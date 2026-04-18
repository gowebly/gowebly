//go:build wireinject

package cmd

import (
	"github.com/google/wire"

	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/gowebly/gowebly/v3/internal/injectors"
)

// Inject initializes the dependency graph using Google Wire.
func Inject() *injectors.Injector {
	wire.Build(attachments.New, config.New, injectors.New)
	return nil
}
