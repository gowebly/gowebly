//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/injector"
)

// inject provides dependency injection process by the "google/wire" package.
func inject(path string) (*injector.Injector, error) {
	wire.Build(attachments.New, config.New, injector.New)
	return &injector.Injector{}, nil
}
