package commands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnknown(t *testing.T) {
	require.NoError(t, Unknown())
}
