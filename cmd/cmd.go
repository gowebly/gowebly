package cmd

import (
	"errors"

	"github.com/gowebly/gowebly/cmd/commands"
	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Run runs cmd commands by the given flags and all dependencies.
func Run(flags []string, di *injector.Injector) error {
	// Check, if flag set is not empty.
	if len(flags) == 0 {
		return errors.New(constants.ErrorRunWithoutFlags)
	}

	helpers.PrintStyled("                          _     _       ", "", "")
	helpers.PrintStyled("   __ _  _____      _____| |__ | |_   _ ", "", "")
	helpers.PrintStyled("  / _` |/ _ \\ \\ /\\ / / _ \\ '_ \\| | | | |", "", "")
	helpers.PrintStyled(" | (_| | (_) \\ V  V /  __/ |_) | | |_| |", "", "")
	helpers.PrintStyled("  \\__, |\\___/ \\_/\\_/ \\___|_.__/|_|\\__, |", "", "")
	helpers.PrintStyled("  |___/                           |___/ ", "", "")

	// Switch between flags or return error.
	switch flags[0] {
	case "create":
		// Check, if flag set include name of a Go backend.
		if len(flags) == 1 {
			// Creating a new project with a built-in net/http package (by default).
			return commands.Create("", di)
		}

		// Creating a new project with the given Go backend.
		return commands.Create(flags[1], di)
	case "add":
		// Check, if flag set include name of a CSS framework.
		if len(flags) == 1 {
			return errors.New(constants.ErrorRunAddCommandWithoutCSSFrameworkName)
		}

		// Adding the given CSS framework to the frontend part of the project.
		return commands.Add(flags[1], di)
	case "run":
		// Running project in a development mode (non-production).
		return commands.Run(di)
	case "build":
		// Building project to production.
		return commands.Build(di)
	default:
		// Returning error message.
		return errors.New(constants.ErrorRunUnknownCommand)
	}
}
