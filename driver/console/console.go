package console

import (
	"log/slog"
	"os"

	"github.com/cyg-pd/go-slogx/driver"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func init() {
	driver.Register("console", New)
}

func New(opts *slog.HandlerOptions, _ map[string]any) slog.Handler {
	w := os.Stdout
	return tint.NewHandler(w, &tint.Options{
		NoColor: !isatty.IsTerminal(w.Fd()),
	})
}
