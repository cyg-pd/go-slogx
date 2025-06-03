package json

import (
	"log/slog"
	"os"

	"github.com/cyg-pd/go-slogx/driver"
)

func init() {
	driver.Register("json", New)
}

func New(opts *slog.HandlerOptions, _ map[string]any) slog.Handler {
	return slog.NewJSONHandler(os.Stdout, opts)
}
