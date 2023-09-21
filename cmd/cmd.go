package cmd

import (
	"errors"

	"github.com/gowebly/gowebly/cmd/commands"
	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/validators"
)

// Run runs cmd commands by the given flags and all dependencies.
func Run(flags []string) error {
	helpers.PrintStyled("                          _     _       ", "", "")
	helpers.PrintStyled("   __ _  _____      _____| |__ | |_   _ ", "", "")
	helpers.PrintStyled("  / _` |/ _ \\ \\ /\\ / / _ \\ '_ \\| | | | |", "", "")
	helpers.PrintStyled(" | (_| | (_) \\ V  V /  __/ |_) | | |_| |", "", "")
	helpers.PrintStyled("  \\__, |\\___/ \\_/\\_/ \\___|_.__/|_|\\__, |", "", "")
	helpers.PrintStyled("  |___/ build amazing Go web apps |___/ ", "", "")

	// Check, if flag set is not empty.
	if len(flags) == 0 {
		return errors.New(constants.ErrorRunWithoutCommand)
	}

	// Inject all dependencies (config, embed files).
	di, err := inject()
	if err != nil {
		return errors.New(constants.ErrorDependencyInjectionNotComplete)
	}

	// Switch between flags or return error.
	switch flags[0] {
	case "init":
		// Init a default YAML config file (.gowebly.yml) in the current folder.
		return commands.Init(di)
	case "create":
		// Validating a config.
		if err = validators.Validate(di.Config); err != nil {
			return err
		}

		// Creating a new project with the given Go backend.
		return commands.Create(di)
	case "run":
		// Validating a config.
		if err = validators.Validate(di.Config); err != nil {
			return err
		}

		// Running project in a development mode (non-production).
		return commands.Run(di)
	case "build":
		// Validating a config.
		if err = validators.Validate(di.Config); err != nil {
			return err
		}

		// Building project to production.
		return commands.Build(di)
	default:
		// Returning error message.
		return errors.New(constants.ErrorRunUnknownCommand)
	}
}
