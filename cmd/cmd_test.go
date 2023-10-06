package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	testCase1 := Run([]string{})
	require.NoError(t, testCase1)

	testCase2 := Run([]string{"test"})
	require.NoError(t, testCase2)
}
