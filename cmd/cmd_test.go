package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	require.NoError(t, Run([]string{}))
	require.NoError(t, Run([]string{"test"}))
}
