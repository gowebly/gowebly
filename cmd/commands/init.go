package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Init runs the 'init' cmd command.
func Init(di *injector.Injector) error {
	// Read the attachments (embedded) config file.
	cfgFile, err := di.Attachments.Configs.ReadFile("configs/default.yml")
	if err != nil {
		return err
	}

	// Create a config file.
	if err := helpers.MakeFile(constants.YAMLConfigFileName, cfgFile); err != nil {
		return err
	}

	helpers.PrintStyled(
		"Successfully created a default config file in the current folder!",
		"success",
		"margin-top",
	)
	helpers.PrintStyled("Next steps:", "", "margin-top-bottom")
	helpers.PrintStyled("Edit config file with your options and parameters", "info", "margin-left")
	helpers.PrintStyled("Run 'gowebly create' command to create a new project", "info", "margin-left")
	helpers.PrintStyled(
		fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
		"warning", "margin-top-bottom",
	)

	return nil
}
