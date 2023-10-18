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
				EmbedFile:  fmt.Sprintf("templates/backend/%s/go.mod.tmpl", di.Config.Backend.GoFramework),
				OutputFile: "go.mod",
				Data:       di.Config.Backend,
			},
			{
				EmbedFile:  fmt.Sprintf("templates/backend/%s/handlers.go.tmpl", di.Config.Backend.GoFramework),
				OutputFile: "handlers.go",
				Data:       nil,
			},
			{
				EmbedFile:  fmt.Sprintf("templates/backend/%s/server.go.tmpl", di.Config.Backend.GoFramework),
				OutputFile: "server.go",
				Data:       di.Config.Backend,
			},
			{
				EmbedFile:  fmt.Sprintf("templates/backend/%s/main.go.tmpl", di.Config.Backend.GoFramework),
				OutputFile: "main.go",
				Data:       di.Config.Backend,
			},
			{
				EmbedFile:  "templates/misc/gitignore.tmpl",
				OutputFile: ".gitignore",
				Data:       di.Config.Frontend,
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
				EmbedFile:  "templates/frontend/main.html",
				OutputFile: filepath.Join("templates", "main.html"),
			},
			{
				EmbedFile:  "templates/frontend/manifest.json",
				OutputFile: filepath.Join("static", "manifest.json"),
			},
			{
				EmbedFile:  fmt.Sprintf("templates/frontend/%s/index.html", di.Config.Frontend.CSSFramework),
				OutputFile: filepath.Join("templates", "pages", "index.html"),
			},
			{
				EmbedFile:  fmt.Sprintf("templates/frontend/%s/assets/styles.css", di.Config.Frontend.CSSFramework),
				OutputFile: filepath.Join("assets", "styles.css"),
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
				EmbedFile:  "static/favicon.png",
				OutputFile: filepath.Join("static", "favicons", "favicon.png"),
			},
			{
				EmbedFile:  "static/favicon.svg",
				OutputFile: filepath.Join("static", "favicons", "favicon.svg"),
			},
			{
				EmbedFile:  "static/favicon.ico",
				OutputFile: filepath.Join("static", "favicons", "favicon.ico"),
			},
			{
				EmbedFile:  "static/logo.svg",
				OutputFile: filepath.Join("static", "images", "logo.svg"),
			},
			{
				EmbedFile:  "static/favicon.png",
				OutputFile: filepath.Join("static", "favicons", "apple-touch-icon.png"),
			},
			{
				EmbedFile:  "static/favicon.svg",
				OutputFile: filepath.Join("static", "favicons", "manifest-touch-icon.svg"),
			},
			{
				EmbedFile:  "static/manifest-desktop-screenshot.jpeg",
				OutputFile: filepath.Join("static", "favicons", "manifest-screenshot-desktop.jpeg"),
			},
			{
				EmbedFile:  "static/manifest-mobile-screenshot.jpeg",
				OutputFile: filepath.Join("static", "favicons", "manifest-screenshot-mobile.jpeg"),
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
				EmbedFile:  fmt.Sprintf("templates/frontend/%s/package.json.tmpl", di.Config.Frontend.CSSFramework),
				OutputFile: "package.json",
				Data:       di.Config.Frontend,
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
					EmbedFile:  "templates/frontend/tailwindcss/postcssrc.tmpl",
					OutputFile: ".postcssrc",
				},
				{
					EmbedFile:  "templates/frontend/tailwindcss/tailwind.config.js",
					OutputFile: "tailwind.config.js",
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
					EmbedFile:  "templates/frontend/unocss/postcssrc.tmpl",
					OutputFile: ".postcssrc",
				},
				{
					EmbedFile:  "templates/frontend/unocss/uno.config.ts",
					OutputFile: "uno.config.ts",
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
				URL: fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HTMXNameOfCDNRepository, di.Config.Frontend.HTMX,
				),
				OutputFile: filepath.Join("static", "htmx.min.js"),
			},
			{
				URL: fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HyperscriptNameOfCDNRepository, di.Config.Frontend.Hyperscript,
				),
				OutputFile: filepath.Join("static", "hyperscript.min.js"),
			},
		},
	); err != nil {
		return err
	}

	// Dependencies part message.
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
				Name:       "go",
				Options:    []string{"mod", "tidy"},
				SkipOutput: true,
				EnvVars:    nil,
			},
			{
				Name:       frontendRuntime,
				Options:    []string{"install"},
				SkipOutput: true,
				EnvVars:    nil,
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
				Name:       frontendRuntime,
				Options:    []string{"run", "build:dev"},
				SkipOutput: true,
				EnvVars:    nil,
			},
		},
	); err != nil {
		return err
	}

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				Text:  "Successfully created a new project in the current folder!",
				State: "success", Style: "margin-top",
			},
			{
				Text: "Project configuration:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  fmt.Sprintf("Backend: %s", di.Config.Backend.GoFramework),
				State: "info", Style: "margin-left",
			},
			{
				Text: fmt.Sprintf(
					"Server port is %d, timeout (in seconds): %d for read, %d for write",
					di.Config.Backend.Port, di.Config.Backend.Timeout.Read, di.Config.Backend.Timeout.Write,
				),
				State: "info", Style: "margin-left-2",
			},
			{
				Text:  fmt.Sprintf("Frontend: %s", di.Config.Frontend.CSSFramework),
				State: "info", Style: "margin-left",
			},
			{
				Text: fmt.Sprintf(
					"htmx '%s', hyperscript '%s'",
					di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
				),
				State: "info", Style: "margin-left-2",
			},
			{
				Text: "Next steps:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  "Design your business logic and future project architecture",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Add your CSS styles to the './assets/styles.css' file",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Add your PWA (Progressive Web App) settings to the './static/manifest.json' file",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Add your HTML templates to the './templates' folder",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Create new handlers for your HTML templates in the 'handlers.go' file",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Run 'gowebly run' command to run your project in a development (non-production) mode",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "Run 'gowebly build' command to build your project for the production",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "To print all commands, just run 'gowebly' without any commands or options",
				State: "warning", Style: "margin-top",
			},
			{
				Text:  fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
				State: "warning", Style: "margin-bottom",
			},
		},
	)

	return nil
}
