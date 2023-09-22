package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'create' cmd command.
func Create(di *injector.Injector) error {
	// Create backend files from templates.
	if err := helpers.GenerateFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedTemplate{
			{
				fmt.Sprintf("templates/backend/%s/go.mod.tmpl", di.Config.Backend.Name),
				"go.mod", "", nil,
			},
			{
				fmt.Sprintf("templates/backend/%s/handlers.go.tmpl", di.Config.Backend.Name),
				"handlers.go", "", nil,
			},
			{
				fmt.Sprintf("templates/backend/%s/main.go.tmpl", di.Config.Backend.Name),
				"main.go", "", di.Config.Backend,
			},
		},
	); err != nil {
		return err
	}

	// Create a new folder(s).
	if err := helpers.MakeFolders("templates/pages", "templates/components/gowebly"); err != nil {
		return err
	}

	// Copy frontend files from the embed file system.
	if err := helpers.CopyFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedFile{
			{
				"templates/frontend/gowebly/layout",
				"templates",
			},
			{
				"templates/frontend/gowebly/components",
				"templates/components/gowebly",
			},
			{
				"templates/misc",
				"templates",
			},
		},
	); err != nil {
		return err
	}

	// Check, if the config file contain a 'css' option in the 'frontend' block.
	if di.Config.Frontend.CSS != nil {
		// If contains, copy CSS framework files from the embed file system.
		if err := helpers.CopyFromEmbedFS(
			di.Attachments.Templates,
			[]helpers.EmbedFile{
				{
					fmt.Sprintf("templates/frontend/%s", di.Config.Frontend.CSS.Framework),
					"templates/pages",
				},
			},
		); err != nil {
			return err
		}
	} else {
		// Else, copy a default frontend files from the embed file system.
		if err := helpers.CopyFromEmbedFS(
			di.Attachments.Templates,
			[]helpers.EmbedFile{
				{
					"templates/frontend/default",
					"templates/pages",
				},
			},
		); err != nil {
			return err
		}
	}

	helpers.PrintStyled(
		"Successfully created a new project in the current folder!",
		"success",
		"margin-top",
	)
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(
		fmt.Sprintf("Backend: %s, on port %d", di.Config.Backend.Name, di.Config.Backend.Port),
		"info", "margin-left",
	)

	if di.Config.Frontend.CSS == nil {
		helpers.PrintStyled("Frontend: default", "info", "margin-left")
	} else {
		helpers.PrintStyled(
			fmt.Sprintf("Frontend: %s", di.Config.Frontend.CSS.Framework),
			"info", "margin-left",
		)
	}

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
