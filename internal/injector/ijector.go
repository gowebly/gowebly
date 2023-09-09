package injector

import (
	"github.com/gowebly/gowebly/embed"
	"github.com/gowebly/gowebly/internal/config"
)

// Injector represents a struct for the dependency injection.
type Injector struct {
	Config     *config.Config
	EmbedFiles *embed.Files
}

// New creates a new injector instance with config.Config and embed.Files.
func New(cfg *config.Config, embedFiles *embed.Files) (*Injector, error) {
	return &Injector{
		Config:     cfg,
		EmbedFiles: embedFiles,
	}, nil
}
