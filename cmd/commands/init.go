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

	// Read the attachments (embedded) config file.
	cfgFile, err := di.Attachments.Configs.ReadFile("configs/default.yml")
	if err != nil {
		return err
	}

	// Create a config file.
	if err := helpers.MakeFiles([]helpers.File{
		{
			constants.YAMLConfigFileName, cfgFile,
		},
	}); err != nil {
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
				fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
				"warning", "margin-top-bottom",
			},
		},
	)

	return nil
}
