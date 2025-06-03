package slogx

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/cyg-pd/go-slogx/driver"
	_ "github.com/cyg-pd/go-slogx/driver/json"
	_ "github.com/cyg-pd/go-slogx/driver/text"
	slogotel "github.com/remychantenay/slog-otel"
)

type config struct {
	Driver  string
	Level   string
	Source  bool
	Options string
}

func New(opts ...option) *slog.Logger {
	var conf config
	conf.Driver = "text"
	conf.Level = "info"
	conf.Options = "{}"

	for _, opt := range opts {
		opt.apply(&conf)
	}

	driverOpts := map[string]any{}
	_ = json.Unmarshal([]byte(conf.Options), &driverOpts)

	h := driver.Get(conf.Driver, &slog.HandlerOptions{
		AddSource: conf.Source,
		Level:     toLevel(conf.Level),
	}, driverOpts)

	h = slogotel.OtelHandler{Next: h}

	return slog.New(h)
}

func toLevel(lvl string) slog.Leveler {
	switch strings.ToLower(lvl) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	}

	return slog.LevelInfo
}
