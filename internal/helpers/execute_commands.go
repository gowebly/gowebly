package helpers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gowebly/gowebly/v2/internal/messages"
)

// Command represents a struct for cmd commands.
type Command struct {
	SkipOutput       bool
	Name             string
	Options, EnvVars []string
}

// Execute executes all commands with (or without) options.
func Execute(commands []Command) error {
	for _, c := range commands {
		// Create a new cmd execution.
		cmd := exec.Command(c.Name, c.Options...)

		// Add environment variables.
		cmd.Env = append(os.Environ(), c.EnvVars...)

		// Check, if command want to show its output (Stdout, Stderr).
		if !c.SkipOutput {
			// Bind Stdout and Stderr to the standard output.
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		// Run command process.
		if err := cmd.Run(); err != nil {
			return fmt.Errorf(messages.ErrorCMDNotExecuteCommand, c.Name, c.Options, err)
		}
	}

	return nil
}
