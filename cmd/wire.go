//go:build wireinject

package cmd

import (
	"github.com/google/wire"

	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/injector"
)

// inject provides dependency injection process by the "google/wire" package.
func inject() (*injector.Injector, error) {
	wire.Build(attachments.New, config.New, injector.New)
	return &injector.Injector{}, nil
}
