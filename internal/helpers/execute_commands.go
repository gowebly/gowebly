package helpers

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/gowebly/gowebly/internal/messages"
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

// ExecuteInGoroutine executes all commands with (or without) options in separate goroutines.
func ExecuteInGoroutine(ctx context.Context, commands []Command, errCh chan<- error) {
	// Create a WaitGroup.
	var wg sync.WaitGroup

	// Start all commands in goroutines.
	for _, c := range commands {
		// Add one to the WaitGroup.
		wg.Add(1)

		// Start a new goroutine.
		go func(c Command) {
			// Notify the main goroutine about the completion of the current goroutine.
			defer func() {
				wg.Done() // remove one from the WaitGroup
				errCh <- nil
			}()

			// Create a new cmd execution.
			cmd := exec.CommandContext(ctx, c.Name, c.Options...)

			// Add environment variables.
			cmd.Env = append(os.Environ(), c.EnvVars...)

			// Check if command wants to show its output (Stdout, Stderr).
			if !c.SkipOutput {
				// Bind Stdout and Stderr to the standard output.
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}

			// Run the current cmd command.
			if err := cmd.Run(); err != nil {
				errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, "execute command", err)
				return
			}
		}(c)
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	// Notify the main goroutine about the completion of the goroutines.
	select {
	case <-ctx.Done():
		errCh <- ctx.Err()
	default:
		errCh <- nil
	}
}
