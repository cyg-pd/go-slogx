package slogx

import (
	"encoding/json"
	"log/slog"
)

type option interface{ apply(*Config) }
type optionFunc func(*Config)

func (fn optionFunc) apply(cfg *Config) { fn(cfg) }

func WithLevel(lvl slog.Leveler) option {
	return optionFunc(func(c *Config) { c.Log.Level = lvl.Level().String() })
}

func WithSource(b bool) option {
	return optionFunc(func(c *Config) { c.Log.Source = b })
}

func WithDriver(driver string, option any) option {
	return optionFunc(func(c *Config) {
		c.Log.Driver = driver
		if b, err := json.Marshal(option); err == nil {
			c.Log.Options = string(b)
		}
	})
}
