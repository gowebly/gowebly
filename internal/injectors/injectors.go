package injectors

import (
	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
)

// Injector provides access to application-wide dependencies.
type Injector struct {
	Config      *config.Config
	Attachments *attachments.Attachments
}

// New creates an Injector with the given configuration and embedded files.
func New(cfg *config.Config, efs *attachments.Attachments) *Injector {
	return &Injector{
		Config:      cfg,
		Attachments: efs,
	}
}
