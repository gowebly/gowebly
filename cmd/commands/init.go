package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// Init runs the 'init' cmd command.
func Init(di *injectors.Injector) error {
	// Header message.
	helpers.PrintStyled(
		"Generating a default config file... Please wait!",
		"wait", "margin-top",
	)

	// Copy a default config file from the embed file system.
	if err := helpers.CopyFilesFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedFile{
			{
				"configs/default.yml",
				constants.YAMLConfigFileName,
			},
		},
	); err != nil {
		return err
	}

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				"Successfully created a default config file ('.gowebly.yml') in the current folder!",
				"success", "margin-top",
			},
			{
				"Next steps:", "", "margin-top-bottom",
			},
			{
				"Edit the '.gowebly.yml' file by adding your own options and parameters to it",
				"info", "margin-left",
			},
			{
				"Run 'gowebly create' command to create a new project",
				"info", "margin-left",
			},
			{
				"To print all commands, just run 'gowebly' without any commands or options",
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
