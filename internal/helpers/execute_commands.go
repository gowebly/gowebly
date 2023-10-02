package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Command represents a struct for cmd commands.
type Command struct {
	SkipOutput bool
	Name       string
	Options    []string
	EnvVars    []string
}

// Execute executes all commands with (or without) options.
func Execute(commands []Command) error {
	for _, c := range commands {
		// Create a new cmd execution.
		cmd := exec.Command(c.Name, c.Options...)

		// Check, if environment variables have elements.
		if len(c.EnvVars) != 0 {
			// Add all environment variables to the current command.
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, c.EnvVars...)
		}

		// Check, if command want to show its output (Stdout, Stderr).
		if !c.SkipOutput {
			// Bind Stdout and Stderr to the standard output.
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		if err := cmd.Start(); err != nil {
			return fmt.Errorf("cmd.Start: %v", err)
		}

		if err := cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				return fmt.Errorf("Exit Status: %d", exiterr.ExitCode())
			} else {
				return fmt.Errorf("cmd.Wait: %v", err)
			}
		}
	}

	return nil
}
