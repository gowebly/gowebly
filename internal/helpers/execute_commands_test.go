package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	require.NoError(t, Execute([]Command{{true, "uname", []string{"-a"}, nil}}))
	require.NoError(
		t,
		Execute([]Command{
			{false, "uname", []string{"-a"}, []string{"GOWEBLY_TEST=true"}},
		}),
	)
	require.Error(
		t,
		Execute([]Command{
			{true, "unknown", []string{"command"}, nil},
		}),
	)
}

func TestExecuteInGoroutine(t *testing.T) {
	require.NoError(
		t,
		ExecuteInGoroutine([]Command{
			{true, "uname", []string{"-a"}, nil},
		}),
	)
	require.NoError(
		t,
		ExecuteInGoroutine([]Command{
			{false, "uname", []string{"-a"}, []string{"GOWEBLY_TEST=true"}},
		}),
	)
	require.Error(
		t,
		ExecuteInGoroutine([]Command{
			{true, "unknown", []string{"command"}, nil},
		}),
	)
}
