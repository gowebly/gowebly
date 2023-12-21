package cmd

import (
	"errors"
	"fmt"

	"github.com/gowebly/gowebly/cmd/commands"
	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/validators"
)

// Run runs cmd commands by the given flags and all dependencies.
func Run(flags []string) error {
	// Logo message.
	helpers.PrintStyled("   ______                       __     __      ", "", "")
	helpers.PrintStyled("  / ____/____  _      __ ___   / /_   / /__  __", "", "")
	helpers.PrintStyled(" / / __ / __ \\| | /| / // _ \\ / __ \\ / // / / /", "", "")
	helpers.PrintStyled("/ /_/ // /_/ /| |/ |/ //  __// /_/ // // /_/ / ", "", "")
	helpers.PrintStyled("\\____/ \\____/ |__/|__/ \\___//_.___//_/ \\__, /  ", "", "")
	helpers.PrintStyled("     CLI to build amazing Go web apps /____/   ", "", "")

	// Inject all dependencies (config, embed files).
	di, err := inject()
	if err != nil {
		return errors.New(constants.ErrorGoweblyDINotComplete)
	}

	// Check, if flag set is not empty.
	if len(flags) == 0 {
		return commands.Unknown()
	}

	// Validate a config.
	if err := validators.Validate(di.Config); err != nil {
		return err
	}

	// Switch between flags or return error.
	switch flags[0] {
	case "init":
		// Check, if flag set more than 1.
		if len(flags) > 1 {
			return fmt.Errorf(constants.ErrorCMDUnknownFlag, flags[0], flags[1])
		}

		// Init a default YAML config file (.gowebly.yml) in the current folder.
		return commands.Init(di)
	case "create":
		// Check, if flag set more than 1.
		if len(flags) > 1 {
			return fmt.Errorf(constants.ErrorCMDUnknownFlag, flags[0], flags[1])
		}

		// Creating a new project with the given Go backend.
		return commands.Create(di)
	case "run":
		// Check, if flag set more than 1.
		if len(flags) > 1 {
			return fmt.Errorf(constants.ErrorCMDUnknownFlag, flags[0], flags[1])
		}

		// Running project in a development mode (non-production).
		return commands.Run(di)
	case "build":
		// Check, if flag set more than 1.
		if len(flags) > 1 {
			if flags[1] != "--skip-docker" {
				return fmt.Errorf(constants.ErrorCMDUnknownFlag, flags[0], flags[1])
			}

			// Building project to production with skipped Docker part.
			return commands.Build(di, "--skip-docker")
		}

		// Building project to production.
		return commands.Build(di, "")
	case "doctor":
		// Check, if flag set more than 1.
		if len(flags) > 1 {
			return fmt.Errorf(constants.ErrorCMDUnknownFlag, flags[0], flags[1])
		}

		// Show all information about user's system.
		return commands.Doctor(di)
	default:
		// Show help message.
		return commands.Unknown()
	}
}
