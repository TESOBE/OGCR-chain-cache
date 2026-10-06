package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/TESOBE/OGCR-chain-cache/internal/obp"
)

// defaultServeInterval is the loop interval when SYNC_INTERVAL_SECONDS is unset.
const defaultServeInterval = 30

// redeclareEvery is how often the Platform App's Scopes are declared again once
// every required one is held. Until then they are declared before every run,
// so the page shows a grant as soon as an administrator makes it.
const redeclareEvery = time.Hour

// historySize is how many past runs the page lists.
const historySize = 20

//go:embed status.html
var statusHTML string

var statusTemplate = template.Must(template.New("status").Parse(statusHTML))

// server runs the mirror on a loop and serves what it has seen.
type server struct {
	rn       *runner
	interval time.Duration
	runNow   chan struct{}
	started  time.Time

	mu           sync.Mutex
	running      bool
	runStarted   time.Time
	nextRun      time.Time
	history      []*Report // newest first
	app          *obp.PlatformApp
	appErr       string
	appCheckedAt time.Time
}

func serve(ctx context.Context, rn *runner, addr string) error {
	if rn.intervalSeconds <= 0 {
		// Recorded in chain_sync_status, so consumers know what counts as stale.
		rn.intervalSeconds = defaultServeInterval
	}
	s := &server{
		rn:       rn,
		interval: time.Duration(rn.intervalSeconds) * time.Second,
		runNow:   make(chan struct{}, 1),
		started:  time.Now(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.handleStatus)
	mux.HandleFunc("GET /health", s.handleHealth)
	mux.HandleFunc("POST /run", s.handleRun)
	httpServer := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 10 * time.Second}

	listenErr := make(chan error, 1)
	go func() {
		slog.Info("status page listening", "url", "http://"+addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			listenErr <- err
		}
	}()

	go s.loop(ctx)

	select {
	case err := <-listenErr:
		return err
	case <-ctx.Done():
	}
	slog.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return httpServer.Shutdown(shutdownCtx)
}

// loop mirrors once per interval, or sooner when asked through runNow.
func (s *server) loop(ctx context.Context) {
	for {
		s.declareIfDue()

		s.mu.Lock()
		s.running, s.runStarted = true, time.Now()
		s.mu.Unlock()

		rep := s.rn.once(ctx)

		s.mu.Lock()
		s.running = false
		s.history = append([]*Report{rep}, s.history...)
		if len(s.history) > historySize {
			s.history = s.history[:historySize]
		}
		s.nextRun = time.Now().Add(s.interval)
		s.mu.Unlock()

		timer := time.NewTimer(s.interval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-s.runNow:
			timer.Stop()
		case <-timer.C:
		}
	}
}

func (s *server) declareIfDue() {
	s.mu.Lock()
	due := s.app == nil || len(s.app.Missing()) > 0 || time.Since(s.appCheckedAt) > redeclareEvery
	s.mu.Unlock()
	if !due {
		return
	}
	app, err := declarePlatformApp(s.rn.client)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.appCheckedAt = time.Now()
	if err != nil {
		s.appErr = err.Error()
		return
	}
	s.app, s.appErr = app, ""
}

func (s *server) handleRun(w http.ResponseWriter, r *http.Request) {
	select {
	case s.runNow <- struct{}{}:
	default: // a run is already queued
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// health is the /health body: enough for a script or the dev-env page to tell
// whether the mirror is working, without parsing the page.
type health struct {
	Status        string     `json:"status"` // starting, ok, partial or no_chain
	Running       bool       `json:"running"`
	LastRunAt     *time.Time `json:"last_run_at,omitempty"`
	HeadBlock     uint64     `json:"head_block,omitempty"`
	ErrorCount    int        `json:"error_count"`
	MissingScopes []string   `json:"missing_scopes"`
	PlatformApp   string     `json:"platform_app_error,omitempty"`
	NextRunAt     *time.Time `json:"next_run_at,omitempty"`
}

func (s *server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	h := health{Status: "starting", Running: s.running, MissingScopes: []string{}, PlatformApp: s.appErr}
	if len(s.history) > 0 {
		last := s.history[0]
		h.Status, h.LastRunAt, h.HeadBlock, h.ErrorCount = last.RunStatus, &last.Finished, last.HeadBlock, len(last.Errors)
	}
	if s.app != nil {
		h.MissingScopes = obp.ScopeNames(s.app.Missing())
	}
	if !s.running && !s.nextRun.IsZero() {
		next := s.nextRun
		h.NextRunAt = &next
	}
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // read-only status, so the dev-env page can read it
	json.NewEncoder(w).Encode(h)
}

// statusView is everything the page template shows, already formatted.
type statusView struct {
	Status, StatusLabel string
	ChainID             uint64
	OBPURL, RPCURL      string
	Space               string
	Interval            string
	Running             bool
	RunningFor          string
	LastRunAgo          string
	LastRunTook         string
	NextRunIn           string
	HeadBlock           uint64
	SyncRecorded        bool
	Mirrors             []MirrorLine
	Errors              []string
	App                 *obp.PlatformApp
	AppErr              string
	AppMissing          int
	History             []historyRow
	Now                 string
}

type historyRow struct {
	At, Status, Took string
	HeadBlock        uint64
	Records, Errors  int
}

var statusLabels = map[string]string{
	"starting":       "Starting",
	"ok":             "OK",
	"partial":        "Partial: some records failed",
	RunStatusNoChain: "Chain unreachable",
}

func (s *server) handleStatus(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	v := statusView{
		Status:   "starting",
		ChainID:  s.rn.chainID,
		OBPURL:   hostOnly(s.rn.obpURL),
		RPCURL:   hostOnly(s.rn.rpcURL),
		Space:    s.rn.client.Space(),
		Interval: s.interval.String(),
		Running:  s.running,
		App:      s.app,
		AppErr:   s.appErr,
		Now:      time.Now().Format("15:04:05"),
	}
	if s.running {
		v.RunningFor = roundDuration(time.Since(s.runStarted))
	} else if !s.nextRun.IsZero() {
		v.NextRunIn = roundDuration(time.Until(s.nextRun))
	}
	if s.app != nil {
		v.AppMissing = len(s.app.Missing())
	}
	if len(s.history) > 0 {
		last := s.history[0]
		v.Status = last.RunStatus
		v.LastRunAgo = roundDuration(time.Since(last.Finished))
		v.LastRunTook = roundDuration(last.Finished.Sub(last.Started))
		v.HeadBlock, v.SyncRecorded = last.HeadBlock, last.SyncRecorded
		v.Mirrors, v.Errors = last.Mirrors, last.Errors
	}
	for _, r := range s.history {
		row := historyRow{At: r.Finished.Format("15:04:05"), Status: r.RunStatus, Took: roundDuration(r.Finished.Sub(r.Started)), HeadBlock: r.HeadBlock, Errors: len(r.Errors)}
		for _, m := range r.Mirrors {
			row.Records += m.Count
		}
		v.History = append(v.History, row)
	}
	s.mu.Unlock()

	v.StatusLabel = statusLabels[v.Status]
	if v.StatusLabel == "" {
		v.StatusLabel = v.Status
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := statusTemplate.Execute(w, v); err != nil {
		slog.Error("render status page", "err", err)
	}
}

// hostOnly drops everything after the host from a URL, since an RPC URL can
// carry an API key in its path or query.
func hostOnly(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "(not a URL)"
	}
	return u.Scheme + "://" + u.Host
}

// roundDuration formats a duration for people: 4s, 2m10s, 3h5m0s.
func roundDuration(d time.Duration) string {
	switch {
	case d < 0:
		return "0s"
	case d < time.Minute:
		return fmt.Sprintf("%ds", int(d.Seconds()))
	case d < time.Hour:
		return d.Round(time.Second).String()
	default:
		return d.Round(time.Minute).String()
	}
}
