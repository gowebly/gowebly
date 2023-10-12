package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKoanfByPath(t *testing.T) {
	t.Run("Existing Config File", func(t *testing.T) {
		koanf, err := newKoanfByPath("internal/configs/default.yml")
		assert.NoError(t, err)
		assert.NotNil(t, koanf)
	})

	t.Run("Default Config File", func(t *testing.T) {
		koanf, err := newKoanfByPath("/path/to/non-existing/config.yml")
		assert.NoError(t, err)
		assert.NotNil(t, koanf)
	})
}
