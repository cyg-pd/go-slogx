package slogx

import (
	"encoding/json"
	"flag"
	"log/slog"
	"strings"

	"github.com/cyg-pd/go-config"
	"github.com/cyg-pd/go-slogx/driver"
	slogotel "github.com/remychantenay/slog-otel"
)

func init() {
	flag.String("log.driver", "stdout", "Log driver")
	flag.String("log.level", "info", "Log level (debug, info, warn, error)")
	flag.String("log.opts", "{}", "Log options (json)")
	flag.Bool("log.source", false, "Log source")
}

type Config struct {
	Log struct {
		Driver  string `mapstructure:"driver"`
		Level   string `mapstructure:"level"`
		Source  bool   `mapstructure:"source"`
		Options string `mapstructure:"opts"`
	} `mapstructure:"log"`
}

func New(opts ...option) *slog.Logger {
	var conf Config
	conf.Log.Driver = "stdout"
	conf.Log.Level = "info"
	conf.Log.Options = "{}"

	_ = config.Unmarshal(&conf)

	for _, opt := range opts {
		opt.apply(&conf)
	}

	driverOpts := map[string]any{}
	_ = json.Unmarshal([]byte(conf.Log.Options), &driverOpts)

	h := driver.Get(conf.Log.Driver, &slog.HandlerOptions{
		AddSource: conf.Log.Source,
		Level:     toLevel(conf.Log.Level),
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
