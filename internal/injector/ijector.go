package injector

import (
	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
)

// Injector represents a struct for the dependency injection.
type Injector struct {
	Config     *config.Config
	EmbedFiles *attachments.Files
}

// New creates a new injector instance with config.Config and embed.Files.
func New(cfg *config.Config, embedFiles *attachments.Files) (*Injector, error) {
	return &Injector{
		Config:     cfg,
		EmbedFiles: embedFiles,
	}, nil
}
