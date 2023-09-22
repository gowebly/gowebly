package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/generators"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'create' cmd command.
func Create(di *injector.Injector) error {
	// Create backend files from templates.
	if err := generators.FilesFromEmbedFS(
		di.Attachments.Templates,
		[]generators.Template{
			{
				fmt.Sprintf("templates/backend/%s/main.go.tmpl", di.Config.Backend.Name),
				"main.go", di.Config.Backend,
			},
			{
				fmt.Sprintf("templates/backend/%s/handlers.go.tmpl", di.Config.Backend.Name),
				"handlers.go", nil,
			},
			{
				fmt.Sprintf("templates/backend/%s/go.mod.tmpl", di.Config.Backend.Name),
				"go.mod", nil,
			},
		},
	); err != nil {
		return err
	}

	helpers.PrintStyled(
		"Successfully created a new project in the current folder!",
		"success",
		"margin-top",
	)
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(fmt.Sprintf(
		"Backend: %s, on port %d", di.Config.Backend.Name, di.Config.Backend.Port),
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
