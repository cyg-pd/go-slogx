package driver

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var outputStore = map[string]func(opts *slog.HandlerOptions, driverOpts map[string]any) slog.Handler{}

func init() {
	Register("stdout", func(opts *slog.HandlerOptions, _ map[string]any) slog.Handler {
		return slog.NewTextHandler(os.Stdout, opts)
	})
	Register("stdout_json", func(opts *slog.HandlerOptions, _ map[string]any) slog.Handler {
		return slog.NewJSONHandler(os.Stdout, opts)
	})
}

func Register(driver string, handler func(opts *slog.HandlerOptions, driverOpts map[string]any) slog.Handler) {
	outputStore[strings.ToLower(driver)] = handler
}

func Get(driver string, opts *slog.HandlerOptions, driverOpts map[string]any) slog.Handler {
	if v, ok := outputStore[strings.ToLower(driver)]; ok {
		return v(opts, driverOpts)
	}

	allow := make([]string, 0, len(outputStore))
	for key := range outputStore {
		allow = append(allow, key)
	}

	panic(fmt.Errorf("invalid log.driver: %s (allowed: %s)", driver, strings.Join(allow, ", ")))
}
