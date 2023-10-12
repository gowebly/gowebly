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
				Text:  "Oops... You've not specified any command to execute, or such a command is not supported!",
				State: "error", Style: "margin-top",
			},
			{
				Text:  "Try to start with these commands:",
				State: "", Style: "margin-top-bottom",
			},
			{
				Text:  "'init' to create a default config file (.gowebly.yml) in the current folder",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "'create' to create a new project (by a default or user config) in the current folder",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "'run' to run your project in a development (non-production) mode",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "'build [OPTION]' to build your project for production and prepare Docker files for deploy",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "option '--skip-docker' to skip generation process for the Docker files",
				State: "info", Style: "margin-left-2",
			},
			{
				Text:  "'doctor' to show all information about your system (for self-debug or issue on GitHub)",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "To print this message, just run 'gowebly' without any commands or options",
				State: "warning", Style: "margin-top",
			},
			{
				Text:  fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
				State: "warning", Style: "margin-bottom",
			},
		},
	)

	return nil
}
