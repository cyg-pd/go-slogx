module github.com/cyg-pd/go-slogx

go 1.24.0

require (
	github.com/cyg-pd/go-otelx v0.0.6
	github.com/lmittmann/tint v1.1.2
	github.com/mattn/go-isatty v0.0.20
	github.com/remychantenay/slog-otel v1.3.4
	go.opentelemetry.io/contrib/bridges/otelslog v0.14.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cyg-pd/go-reflectx v0.0.1 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/log v0.15.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/sdk v1.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
)

retract (
	v0.0.3 // This version support flag config but drop support in 0.0.4. Please use version 0.0.4 instead.
	v0.0.2 // This version support flag config but drop support in 0.0.4. Please use version 0.0.4 instead.
	v0.0.1 // This version support flag config but drop support in 0.0.4. Please use version 0.0.4 instead.
)
