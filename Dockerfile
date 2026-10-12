FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Two binaries: the cacher service and the setup-entity tool. Both carry the
# entities/*.json definitions built in; the cacher applies them on start.
RUN CGO_ENABLED=0 go build -o /out/cacher ./cmd/cacher \
 && CGO_ENABLED=0 go build -o /out/setup-entity ./cmd/setup-entity

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/cacher /usr/local/bin/cacher
COPY --from=builder /out/setup-entity /usr/local/bin/setup-entity
# Default to the cacher service in serve mode: it mirrors every
# SYNC_INTERVAL_SECONDS (default 30), re-declares its Platform App Scopes, and
# serves the status page on SERVE_ADDR. Inside the container that has to be all
# interfaces; publish the port only to localhost or a private network, since the
# page needs no login and has a Run now button. For a single run, override the
# command with `cacher`; for the entity tool, with `setup-entity`.
ENV SERVE_ADDR=0.0.0.0:8766
EXPOSE 8766
CMD ["cacher", "-serve"]
