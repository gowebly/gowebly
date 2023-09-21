package helpers

import (
	"os"
	"os/exec"
)

// Execute executes the given command with (or without) options.
func Execute(command string, options ...string) error {
	// Create a new cmd execution.
	cmd := exec.Command(command, options...)

	// Bind Stdout and Stderr.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
