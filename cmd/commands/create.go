package commands

import (
	"fmt"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// Create runs the 'create' cmd command.
func Create(di *injectors.Injector) error {
	// Backend part message.
	helpers.PrintStyled(
		"Preparing the backend part... Please wait!",
		"wait", "margin-top",
	)

	// Create a new folder(s).
	if err := helpers.MakeFolders("assets", "static/favicons", "static/images", "templates/pages"); err != nil {
		return err
	}

	// Create backend and misc files from templates.
	if err := helpers.GenerateFilesByTemplateFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedTemplate{
			{
				fmt.Sprintf("templates/backend/%s/go.mod.tmpl", di.Config.Backend.GoFramework),
				"go.mod",
				di.Config.Backend,
			},
			{
				fmt.Sprintf("templates/backend/%s/handlers.go.tmpl", di.Config.Backend.GoFramework),
				"handlers.go",
				nil,
			},
			{
				fmt.Sprintf("templates/backend/%s/server.go.tmpl", di.Config.Backend.GoFramework),
				"server.go",
				di.Config.Backend,
			},
			{
				fmt.Sprintf("templates/backend/%s/main.go.tmpl", di.Config.Backend.GoFramework),
				"main.go",
				di.Config.Backend,
			},
			{
				"templates/misc/gitignore.tmpl",
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
				"templates/frontend/main.html",
				filepath.Join("templates", "main.html"),
			},
			{
				fmt.Sprintf("templates/frontend/%s/index.html", di.Config.Frontend.CSSFramework),
				filepath.Join("templates", "pages", "index.html"),
			},
			{
				fmt.Sprintf("templates/frontend/%s/assets/styles.css", di.Config.Frontend.CSSFramework),
				filepath.Join("assets", "styles.css"),
			},
		},
	); err != nil {
		return err
	}

	// Copy static files from the embed file system.
	if err := helpers.CopyFilesFromEmbedFS(
		di.Attachments.Static,
		[]helpers.EmbedFile{
			{
				"static/apple-touch-icon.png",
				filepath.Join("static", "favicons", "apple-touch-icon.png"),
			},
			{
				"static/favicon.png",
				filepath.Join("static", "favicons", "favicon.png"),
			},
			{
				"static/favicon.svg",
				filepath.Join("static", "favicons", "favicon.svg"),
			},
			{
				"static/favicon.ico",
				filepath.Join("static", "favicons", "favicon.ico"),
			},
			{
				"static/logo.svg",
				filepath.Join("static", "images", "logo.svg"),
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
				fmt.Sprintf("templates/frontend/%s/package.json.tmpl", di.Config.Frontend.CSSFramework),
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
					"templates/frontend/tailwindcss/postcssrc.tmpl",
					".postcssrc",
				},
				{
					"templates/frontend/tailwindcss/tailwind.config.js",
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
					"templates/frontend/unocss/postcssrc.tmpl",
					".postcssrc",
				},
				{
					"templates/frontend/unocss/uno.config.ts",
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

	// Set the default JavaScript runtime environment.
	frontendRuntime := "npm"

	// Check, if the runtime of the frontend part is switched.
	if di.Config.Frontend.RuntimeEnvironment == "bun" {
		frontendRuntime = "bun"
	}

	// Execute system commands.
	if err := helpers.ExecuteInGoroutine(
		[]helpers.Command{
			{
				true,
				"go",
				[]string{"mod", "tidy"},
				nil,
			},
			{
				true,
				frontendRuntime,
				[]string{"install"},
				nil,
			},
		},
	); err != nil {
		return err
	}

	// Frontend build message.
	helpers.PrintStyled(
		"Building a CSS bundle in the development (non-production) mode for the frontend part... Please wait!",
		"wait", "",
	)

	// Execute system commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				true,
				frontendRuntime,
				[]string{"run", "build:dev"},
				nil,
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
				"To print all commands, just run 'gowebly' without any commands or options",
				"warning", "margin-top",
			},
			{
				fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
				"warning", "margin-bottom",
			},
		},
	)

	return nil
}
