package cmd

import (
	"github.com/gowebly/gowebly/v3/cmd/commands"
)

// Run routes CLI flags to the appropriate command handler.
func Run(flags []string) error {
	if len(flags) == 0 {
		return commands.Unknown()
	}

	switch flags[0] {
	case "create":
		di := Inject()
		return commands.Create(di)
	case "run":
		return commands.Run()
	case "doctor":
		return commands.Doctor()
	default:
		return commands.Unknown()
	}
}
