package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_inject(t *testing.T) {
	_, err := inject()
	require.NoError(t, err)
}
