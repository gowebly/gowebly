package helpers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gowebly/gowebly/v3/internal/messages"
)

// Command represents an external command to be executed.
type Command struct {
	SkipOutput       bool
	Name             string
	Options, EnvVars []string
}

// Execute runs all commands in sequence, stopping on first error.
func Execute(commands []Command) error {
	for _, c := range commands {
		cmd := exec.Command(c.Name, c.Options...)
		cmd.Env = append(os.Environ(), c.EnvVars...)

		if !c.SkipOutput {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		if err := cmd.Run(); err != nil {
			return fmt.Errorf(messages.ErrorCMDNotExecuteCommand, c.Name, c.Options, err)
		}
	}

	return nil
}
