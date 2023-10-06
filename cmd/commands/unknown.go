package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
)

// Unknown runs an unknown cmd command or nothing.
func Unknown() error {
	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				"Oops... You've not specified any command to execute, or such a command is not supported!",
				"error", "margin-top",
			},
			{
				"Try to start with these commands:",
				"", "margin-top-bottom",
			},
			{
				"'init' to create a default config file (.gowebly.yml) in the current folder",
				"info", "margin-left",
			},
			{
				"'create' to create a new project (by a default or user config) in the current folder",
				"info", "margin-left",
			},
			{
				"'run' to run your project in a development (non-production) mode",
				"info", "margin-left",
			},
			{
				"'build [OPTION]' to build your project for production and prepare Docker files for deploy",
				"info", "margin-left",
			},
			{
				"option '--skip-docker' to skip generation process for the Docker files",
				"info", "margin-left-2",
			},
			{
				"'doctor' to show all information about your system (for self-debug or issue on GitHub)",
				"info", "margin-left",
			},
			{
				"To print this message, just run 'gowebly' without any commands or options",
				"warning", "margin-top",
			},
			{
				fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
				"warning", "margin-bottom",
			},
		},
	)

	return nil
}
