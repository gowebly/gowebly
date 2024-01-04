package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T) {
	assert.Nil(t, Run(nil), nil)
	assert.Nil(t, Run([]string{"unknown"}), nil)
	assert.Nil(t, Run([]string{"doctor"}), nil)
	assert.Error(t, Run([]string{"create"}))
	assert.Error(t, Run([]string{"run"}))
}
