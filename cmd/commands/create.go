package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'create' cmd command.
func Create(di *injector.Injector) error {
	helpers.PrintStyled(
		"Successfully created a new project in the current folder!",
		"success",
		"margin-top",
	)
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(fmt.Sprintf(
		"Go backend: %s, on port %d", di.Config.Backend.Name, di.Config.Backend.Port),
		"info", "margin-left",
	)
	helpers.PrintStyled("Next steps:", "", "margin-top-bottom")
	helpers.PrintStyled("Add your business logic to the project files", "info", "margin-left")
	helpers.PrintStyled(
		"Run 'gowebly run' command to run your project in a developer mode",
		"info", "margin-left",
	)
	helpers.PrintStyled(
		"Run 'gowebly build' command to build your project for the production",
		"info", "margin-left",
	)
	helpers.PrintStyled(fmt.Sprintf(
		"For more information, see %s", constants.LinkToCompleteUserGuide),
		"warning", "margin-top-bottom",
	)

	return nil
}
