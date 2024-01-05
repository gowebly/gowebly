package commands

import (
	"testing"

	"github.com/gowebly/gowebly/v2/internal/attachments"
	"github.com/gowebly/gowebly/v2/internal/config"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	di := &injectors.Injector{
		Config:      &config.Config{Backend: &config.Backend{}, Frontend: &config.Frontend{}},
		Attachments: &attachments.Attachments{},
	}
	assert.Error(t, Create(di))
}
