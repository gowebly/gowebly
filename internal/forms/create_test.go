package forms

import (
	"testing"

	"github.com/gowebly/gowebly/internal/attachments"
	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/stretchr/testify/assert"
)

func Test_htmxForm(t *testing.T) {
	di := &injectors.Injector{
		Config:      &config.Config{Backend: &config.Backend{}, Frontend: &config.Frontend{}, Tools: &config.Tools{}},
		Attachments: &attachments.Attachments{},
	}
	assert.Error(t, projectSettingsForm(di))
}
