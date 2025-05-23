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

func WithDriver(driver string) option {
	return optionFunc(func(c *Config) { c.Log.Driver = driver })
}

func WithDriverOptions(option any) option {
	return optionFunc(func(c *Config) {
		if b, err := json.Marshal(option); err == nil {
			c.Log.Options = string(b)
		}
	})
}
