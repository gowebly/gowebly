package commands

import (
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
	if err = helpers.MakeFile(constants.YAMLConfigFileName, cfgFile); err != nil {
		return err
	}

	helpers.PrintStyled(
		"Successfully created a default config file in the current folder!",
		"success",
		"margin-top-bottom",
	)

	return nil
}
