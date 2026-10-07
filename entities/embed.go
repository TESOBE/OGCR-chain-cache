// Package entities holds the OBP dynamic entity definitions this tool writes
// to, built into the binaries so the cacher can apply them on start without a
// copy of this directory next to it.
package entities

import "embed"

//go:embed *.json
var FS embed.FS
