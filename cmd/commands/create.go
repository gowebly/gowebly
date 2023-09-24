package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'create' cmd command.
func Create(di *injector.Injector) error {
	// Header message.
	helpers.PrintStyled(
		"Downloading and preparing a minified versions of the frontend part... Please wait!",
		"info", "margin-top",
	)

	// Create a new folder(s).
	if err := helpers.MakeFolders("static", "templates/pages", "templates/components/gowebly"); err != nil {
		return err
	}

	// Create backend and misc files from templates.
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
				fmt.Sprintf("templates/backend/%s/server.go.tmpl", di.Config.Backend.Name),
				"server.go", "", nil,
			},
			{
				fmt.Sprintf("templates/backend/%s/main.go.tmpl", di.Config.Backend.Name),
				"main.go", "", di.Config.Backend,
			},
			{
				"templates/misc/gitignore.tmpl",
				".gitignore", "", nil,
			},
			{
				"templates/misc/env.tmpl",
				".env", "", di.Config.Backend,
			},
		},
	); err != nil {
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
				"templates/static",
				"static",
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

	// Download minified version of the htmx and hyperscript JS files from CND.
	if err := helpers.DownloadFiles(
		[]helpers.Download{
			{
				fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HTMXNameOfCDNRepository, di.Config.Frontend.HTMX,
				),
				"htmx.min.js",
				"static",
			},
			{
				fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HyperscriptNameOfCDNRepository, di.Config.Frontend.Hyperscript,
				),
				"hyperscript.min.js",
				"static",
			},
		},
	); err != nil {
		return err
	}

	// Success message.
	helpers.PrintStyled(
		"Successfully created a new project in the current folder!",
		"success",
		"margin-top",
	)

	// Project config message.
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(
		fmt.Sprintf("Backend: %s", di.Config.Backend.Name),
		"info", "margin-left",
	)
	helpers.PrintStyled(
		fmt.Sprintf(
			"Server port is %d, timeout (in seconds): %d for read, %d for write",
			di.Config.Backend.Port, di.Config.Backend.Timeout.Read, di.Config.Backend.Timeout.Write,
		),
		"info", "margin-left-2",
	)

	if di.Config.Frontend.CSS == nil {
		helpers.PrintStyled("Frontend: default", "info", "margin-left")
	} else {
		helpers.PrintStyled(
			fmt.Sprintf("Frontend: %s ('%s')", di.Config.Frontend.CSS.Framework, di.Config.Frontend.CSS.Version),
			"info", "margin-left",
		)
	}

	helpers.PrintStyled(
		fmt.Sprintf(
			"htmx ('%s'), hyperscript ('%s')",
			di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
		),
		"info", "margin-left-2",
	)

	// Next steps message.
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

	// Footer message.
	helpers.PrintStyled(fmt.Sprintf(
		"For more information, see %s", constants.LinkToCompleteUserGuide),
		"warning", "margin-top-bottom",
	)

	return nil
}
