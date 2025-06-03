package slogx

import (
	"encoding/json"
	"log/slog"
)

type option interface{ apply(*config) }
type optionFunc func(*config)

func (fn optionFunc) apply(cfg *config) { fn(cfg) }

func WithLevel(lvl string) option {
	return optionFunc(func(c *config) { c.Level = lvl })
}

func WithSlogLevel(lvl slog.Leveler) option {
	return optionFunc(func(c *config) { c.Level = lvl.Level().String() })
}

func WithSource(b bool) option {
	return optionFunc(func(c *config) { c.Source = b })
}

func WithDriver(driver string, option any) option {
	return optionFunc(func(c *config) {
		c.Driver = driver
		if b, err := json.Marshal(option); err == nil {
			c.Options = string(b)
		}
	})
}
