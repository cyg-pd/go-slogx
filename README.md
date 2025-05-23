# go-slogx

[![tag](https://img.shields.io/github/tag/cyg-pd/go-slogx.svg)](https://github.com/cyg-pd/go-slogx/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/cyg-pd/go-slogx?status.svg)](https://pkg.go.dev/github.com/cyg-pd/go-slogx)
![Build Status](https://github.com/cyg-pd/go-slogx/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/cyg-pd/go-slogx)](https://goreportcard.com/report/github.com/cyg-pd/go-slogx)
[![Coverage](https://img.shields.io/codecov/c/github/cyg-pd/go-slogx)](https://codecov.io/gh/cyg-pd/go-slogx)
[![Contributors](https://img.shields.io/github/contributors/cyg-pd/go-slogx)](https://github.com/cyg-pd/go-slogx/graphs/contributors)
[![License](https://img.shields.io/github/license/cyg-pd/go-slogx)](./LICENSE)

## 🚀 Install

```sh
go get github.com/cyg-pd/go-slogx@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

This library has no dependencies outside the Go standard library.

## 💡 Usage

You can import `go-slogx` using:

```go
package main

import (
	"context"
	"log/slog"

	"github.com/cyg-pd/go-config"
	"github.com/cyg-pd/go-slogx"
	"github.com/cyg-pd/go-otelx"
	_ "github.com/cyg-pd/go-otelx/autoconf"
)

var tracer = otelx.Tracer()

func init() {
	config.Parse()
	slog.SetDefault(slogx.New())
}

func main() {
	ctx := context.Background()
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	slog.InfoContext(ctx, "hi") // time=2025-05-05T00:00:00.000+08:00 level=INFO msg=hi trace_id=e603b53b1ae2f90397dc8768301fa857 span_id=7fd54f3aafe4a95a
}
```
