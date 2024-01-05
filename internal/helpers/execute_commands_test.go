package helpers

import (
	"testing"
)

func TestExecute(t *testing.T) {
	// Test case 1: Test executing a single command without options.
	t.Run("Test executing a single command without options", func(t *testing.T) {
		commands := []Command{
			{
				Name:       "echo",
				Options:    []string{},
				EnvVars:    []string{},
				SkipOutput: false,
			},
		}

		err := Execute(commands)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})
}
