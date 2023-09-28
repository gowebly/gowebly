package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	_, err := New()
	require.NoError(t, err, &Config{})
}
