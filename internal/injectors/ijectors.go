package injectors

import (
	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
)

// Injector represents a struct for the dependency injection.
type Injector struct {
	Config      *config.Config
	Attachments *attachments.Files
}

// New creates a new injector instance with config.Config and attachments.Files.
func New(cfg *config.Config, efs *attachments.Files) (*Injector, error) {
	return &Injector{
		Config:      cfg,
		Attachments: efs,
	}, nil
}
