package text

import (
	"log/slog"
	"os"

	"github.com/cyg-pd/go-slogx/driver"
)

func init() {
	driver.Register("text", New)
}

func New(opts *slog.HandlerOptions, _ map[string]any) slog.Handler {
	return slog.NewTextHandler(os.Stdout, opts)
}
