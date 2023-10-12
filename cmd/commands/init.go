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
		di.Attachments.Configs,
		[]helpers.EmbedFile{
			{
				EmbedFile:  "configs/default.yml",
				OutputFile: constants.YAMLConfigFileName,
			},
		},
	); err != nil {
		return err
	}

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				Text:  "Successfully created a default config file ('.gowebly.yml') in the current folder!",
				State: "success", Style: "margin-top",
			},
			{
				Text: "Next steps:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  "Edit the '.gowebly.yml' file by adding your own options and parameters to it",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Run 'gowebly create' command to create a new project",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "To print all commands, just run 'gowebly' without any commands or options",
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
