//go:build wireinject

package cmd

import (
	"github.com/google/wire"

	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/gowebly/gowebly/v3/internal/injectors"
)

// Inject provides the dependency injection process by the "google/wire" package.
func Inject() (*injectors.Injector, error) {
	wire.Build(attachments.New, config.New, injectors.New)
	return &injectors.Injector{}, nil
}
