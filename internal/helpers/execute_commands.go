package helpers

import (
	"context"
	"fmt"
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

		// Run command process.
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("cmd: can't execute command '%s %s'", c.Name, c.Options)
		}
	}

	return nil
}

// ExecuteInGoroutine executes all commands with (or without) options in a
// separated goroutines.
func ExecuteInGoroutine(commands []Command) error {
	// Create context and cancel function.
	ctx, cancel := context.WithCancel(context.Background())

	// Create error channel.
	errChan := make(chan error, len(commands))

	for _, c := range commands {
		// Run command process in goroutine.
		go func(c Command) {
			// Notify the main goroutine about the completion of the current goroutine.
			defer func() { errChan <- nil }()

			// Create a new cmd execution.
			cmd := exec.CommandContext(ctx, c.Name, c.Options...)

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

			// Run the current cmd command.
			if err := cmd.Run(); err != nil {
				errChan <- err
				return
			}
		}(c)
	}

	// Wait for all goroutines to complete.
	for range commands {
		// Cancel the context if an error occurred.
		if err := <-errChan; err != nil {
			cancel()
			return err
		}
	}

	// Cancel the context if the process is successes.
	cancel()

	return nil
}
