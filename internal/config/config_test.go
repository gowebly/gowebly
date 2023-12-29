package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	assert.Equal(t, "github.com/user/project", c.Backend.ModuleName)
	assert.Equal(t, "", c.Backend.GoFramework)
	assert.Equal(t, "7000", c.Backend.Port)
	assert.Equal(t, false, c.Tools.IsUseTempl)
	assert.Equal(t, "project", c.Frontend.PackageName)
	assert.Equal(t, "htmx", c.Frontend.ReactiveLibrary)
	assert.Equal(t, "", c.Frontend.CSSFramework)
	assert.Equal(t, true, c.Tools.IsUseBun)
}
