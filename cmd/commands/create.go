package commands

import (
	"fmt"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'create' cmd command.
func Create(di *injector.Injector) error {
	// Backend part message.
	helpers.PrintStyled(
		"Preparing the backend part... Please wait!",
		"info", "margin-top",
	)

	// Create a new folder(s).
	if err := helpers.MakeFolders("assets", "static", "templates/pages"); err != nil {
		return err
	}

	// Create backend and misc files from templates.
	if err := helpers.GenerateFilesByTemplateFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedTemplate{
			{
				filepath.Join("templates", "backend", di.Config.Backend.Name, "go.mod.tmpl"),
				"go.mod",
				nil,
			},
			{
				filepath.Join("templates", "backend", di.Config.Backend.Name, "handlers.go.tmpl"),
				"handlers.go",
				nil,
			},
			{
				filepath.Join("templates", "backend", di.Config.Backend.Name, "server.go.tmpl"),
				"server.go",
				nil,
			},
			{
				filepath.Join("templates", "backend", di.Config.Backend.Name, "main.go.tmpl"),
				"main.go",
				di.Config.Backend,
			},
			{
				filepath.Join("templates", "misc", "gitignore.tmpl"),
				".gitignore",
				di.Config.Frontend,
			},
			{
				filepath.Join("templates", "misc", "env.tmpl"),
				".env",
				di.Config.Backend,
			},
		},
	); err != nil {
		return err
	}

	// Copy frontend files from the embed file system.
	if err := helpers.CopyFilesFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedFile{
			{
				filepath.Join("templates", "frontend", "main.html"),
				filepath.Join("templates", "main.html"),
			},
			{
				filepath.Join("templates", "frontend", di.Config.Frontend.CSSFramework, "index.html"),
				filepath.Join("templates", "pages", "index.html"),
			},
			{
				filepath.Join("templates", "frontend", di.Config.Frontend.CSSFramework, "assets", "styles.css"),
				filepath.Join("assets", "styles.css"),
			},
			{
				filepath.Join("templates", "frontend", di.Config.Frontend.CSSFramework, "package.json"),
				"package.json",
			},
			{
				filepath.Join("templates", "static", "favicon.ico"),
				filepath.Join("static", "favicon.ico"),
			},
		},
	); err != nil {
		return err
	}

	// Copy CSS framework specific files from the embed file system.
	switch di.Config.Frontend.CSSFramework {
	case "tailwindcss":
		// Tailwind CSS files.
		if err := helpers.CopyFilesFromEmbedFS(
			di.Attachments.Templates,
			[]helpers.EmbedFile{
				{
					filepath.Join("templates", "frontend", "tailwindcss", "postcss.config.js"),
					"postcss.config.js",
				},
				{
					filepath.Join("templates", "frontend", "tailwindcss", "tailwind.config.js"),
					"tailwind.config.js",
				},
			},
		); err != nil {
			return err
		}
	case "unocss":
		// UnoCSS files.
	}

	// Frontend part message.
	helpers.PrintStyled(
		"Downloading and preparing the frontend part... Please wait!",
		"info", "",
	)

	// Download minified version of the htmx and hyperscript JS files from CND.
	if err := helpers.DownloadFiles(
		[]helpers.Download{
			{
				fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HTMXNameOfCDNRepository, di.Config.Frontend.HTMX,
				),
				filepath.Join("static", "htmx.min.js"),
			},
			{
				fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HyperscriptNameOfCDNRepository, di.Config.Frontend.Hyperscript,
				),
				filepath.Join("static", "hyperscript.min.js"),
			},
		},
	); err != nil {
		return err
	}

	// Execute system commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				true, "npm", []string{"install"},
			},
			{
				true, "npm", []string{"run", "build:dev"},
			},
			{
				true, "go", []string{"mod", "tidy"},
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
	helpers.PrintStyled(
		fmt.Sprintf("Frontend: %s", di.Config.Frontend.CSSFramework),
		"info", "margin-left",
	)
	helpers.PrintStyled(
		fmt.Sprintf(
			"htmx '%s', hyperscript '%s'",
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
