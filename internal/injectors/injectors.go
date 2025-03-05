package injectors

import (
	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
)

// Injector represents a struct for the dependency injection.
type Injector struct {
	Config      *config.Config
	Attachments *attachments.Attachments
}

// New creates a new injector instance with config.Config and attachments.Attachments.
func New(cfg *config.Config, efs *attachments.Attachments) (*Injector, error) {
	return &Injector{
		Config:      cfg,
		Attachments: efs,
	}, nil
}
