package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoctor(t *testing.T) {
	assert.Nil(t, Doctor(), nil)
}
