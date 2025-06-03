package otel

import (
	"log/slog"

	"github.com/cyg-pd/go-otelx"
	"github.com/cyg-pd/go-slogx/driver"
	"go.opentelemetry.io/contrib/bridges/otelslog"
)

func init() {
	driver.Register("otel", New)
}

func New(_ *slog.HandlerOptions, _ map[string]any) slog.Handler {
	return otelslog.NewHandler(otelx.ServiceName())
}
