package cmd

import (
	"fmt"

	"github.com/gowebly/gowebly/cmd/commands"
	"github.com/gowebly/gowebly/internal/messages"
)

// Run runs cmd commands by the given flags and all dependencies.
func Run(flags []string) error {
	// Check, if flag set is not empty.
	if len(flags) == 0 {
		// Run 'unknown' command.
		return commands.Unknown()
	}

	// Switch between flags or return error.
	switch flags[0] {
	case "create":
		// Inject all dependencies (config, embed files).
		di, err := Inject()
		if err != nil {
			return fmt.Errorf(messages.ErrorGoweblyDINotComplete, err)
		}

		// Run 'create' command.
		return commands.Create(di)
	case "doctor":
		// Run 'doctor' command.
		return commands.Doctor()
	default:
		// Run 'unknown' command by default.
		return commands.Unknown()
	}
}
