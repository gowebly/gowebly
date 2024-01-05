package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknown(t *testing.T) {
	assert.Nil(t, Unknown(), nil)
}
