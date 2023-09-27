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
		"wait", "margin-top",
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
				filepath.Join("templates", "backend", di.Config.Backend.GoFramework, "go.mod.tmpl"),
				"go.mod",
				di.Config.Backend,
			},
			{
				filepath.Join("templates", "backend", di.Config.Backend.GoFramework, "handlers.go.tmpl"),
				"handlers.go",
				nil,
			},
			{
				filepath.Join("templates", "backend", di.Config.Backend.GoFramework, "server.go.tmpl"),
				"server.go",
				di.Config.Backend,
			},
			{
				filepath.Join("templates", "backend", di.Config.Backend.GoFramework, "main.go.tmpl"),
				"main.go",
				di.Config.Backend,
			},
			{
				filepath.Join("templates", "misc", "gitignore.tmpl"),
				".gitignore",
				di.Config.Frontend,
			},
		},
	); err != nil {
		return err
	}

	// Frontend part message.
	helpers.PrintStyled(
		"Downloading and preparing the frontend part... Please wait!",
		"wait", "",
	)

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
				filepath.Join(
					"templates", "frontend", di.Config.Frontend.CSSFramework, "assets", "styles.css",
				),
				filepath.Join("assets", "styles.css"),
			},
			{
				filepath.Join("templates", "static", "favicon.ico"),
				filepath.Join("static", "favicon.ico"),
			},
		},
	); err != nil {
		return err
	}

	// Create frontend files from templates.
	if err := helpers.GenerateFilesByTemplateFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedTemplate{
			{
				filepath.Join(
					"templates", "frontend", di.Config.Frontend.CSSFramework, "package.json.tmpl",
				),
				"package.json",
				di.Config.Frontend,
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
		if err := helpers.CopyFilesFromEmbedFS(
			di.Attachments.Templates,
			[]helpers.EmbedFile{
				{
					filepath.Join("templates", "frontend", "unocss", "postcss.config.js"),
					"postcss.config.js",
				},
				{
					filepath.Join("templates", "frontend", "unocss", "uno.config.ts"),
					"uno.config.ts",
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

	// Frontend part message.
	helpers.PrintStyled(
		"Fetching dependencies of the backend and frontend parts... Please wait!",
		"wait", "",
	)

	// Execute system commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				true, "npm", []string{"install"}, nil,
			},
			{
				true, "npm", []string{"run", "build:dev"}, nil,
			},
			{
				true, "go", []string{"mod", "tidy"}, nil,
			},
		},
	); err != nil {
		return err
	}

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				"Successfully created a new project in the current folder!",
				"success", "margin-top",
			},
			{
				"Project configuration:", "", "margin-top-bottom",
			},
			{
				fmt.Sprintf("Backend: %s", di.Config.Backend.GoFramework),
				"info", "margin-left",
			},
			{
				fmt.Sprintf(
					"Server port is %d, timeout (in seconds): %d for read, %d for write",
					di.Config.Backend.Port, di.Config.Backend.Timeout.Read, di.Config.Backend.Timeout.Write,
				),
				"info", "margin-left-2",
			},
			{
				fmt.Sprintf("Frontend: %s", di.Config.Frontend.CSSFramework),
				"info", "margin-left",
			},
			{
				fmt.Sprintf(
					"htmx '%s', hyperscript '%s'",
					di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
				),
				"info", "margin-left-2",
			},
			{
				"Next steps:", "", "margin-top-bottom",
			},
			{
				"Design your business logic and future project architecture",
				"info", "margin-left",
			},
			{
				"Add your CSS styles to the './assets/styles.css' file",
				"info", "margin-left",
			},
			{
				"Add your HTML templates to the './templates' folder",
				"info", "margin-left",
			},
			{
				"Create new handlers for your HTML templates in the 'handlers.go' file",
				"info", "margin-left",
			},
			{
				"Run 'gowebly run' command to run your project in a development (non-production) mode",
				"info", "margin-left",
			},
			{
				"Run 'gowebly build' command to build your project for the production",
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
