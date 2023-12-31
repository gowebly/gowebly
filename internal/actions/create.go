package actions

import (
	"context"
	"fmt"

	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
)

type action struct {
	name string
	fn   func(*injectors.Injector) error
}

// CreateProjectAction creates a new project.
func CreateProjectAction(ctx context.Context, cancel context.CancelFunc, di *injectors.Injector, errCh chan<- error) {
	// Define the list of actions to perform
	actions := []action{
		{"create project folders", createProjectFolders},
		{"copy static files", copyStaticFiles},
		{"generate backend files", createBackendFiles},
		{"generate frontend files", createFrontendFiles},
		{"install project dependencies", installDependencies},
	}

	// Execute each action in the list
	for _, action := range actions {
		if err := action.fn(di); err != nil {
			cancel()
			errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, action.name, err)
			return
		}
	}

	// Check context.
	if ctx.Err() != nil {
		cancel()
		errCh <- ctx.Err()
		return
	}

	cancel()
	errCh <- nil
}

// createProjectFolders creates project folders.
func createProjectFolders(di *injectors.Injector) error {
	// Define a slice of folder paths.
	folders := []string{"static", "web/src"}

	// Check if HTMX is used.
	if di.Config.Frontend.IsUseHTMX {
		// Add templates folders.
		folders = append(folders, "templates/pages")
	}

	return helpers.MakeFolders(folders...)
}

// copyStaticFiles copies static files.
func copyStaticFiles(di *injectors.Injector) error {
	// Define the files to be copied.
	files := []helpers.EmbedFile{
		{
			EmbedFile:  "static/favicon.ico",
			OutputFile: "static/favicon.ico",
		},
		{
			EmbedFile:  "static/favicon.png",
			OutputFile: "static/favicon.png",
		},
		{
			EmbedFile:  "static/favicon.svg",
			OutputFile: "static/favicon.svg",
		},
		{
			EmbedFile:  "static/logo.svg",
			OutputFile: "static/logo.svg",
		},
		{
			EmbedFile:  "static/manifest-desktop-screenshot.jpeg",
			OutputFile: "static/manifest-desktop-screenshot.jpeg",
		},
		{
			EmbedFile:  "static/manifest-mobile-screenshot.jpeg",
			OutputFile: "static/manifest-mobile-screenshot.jpeg",
		},
		{
			EmbedFile:  "static/manifest.json",
			OutputFile: "static/manifest.json",
		},
	}

	return helpers.CopyFilesFromEmbedFS(di.Attachments.Static, files)
}

// createBackendFiles creates backend files.
func createBackendFiles(di *injectors.Injector) error {
	// Define embed templates.
	templates := []helpers.EmbedTemplate{
		// Backend files.
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/go.mod.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "go.mod",
			Data:       di.Config,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/handlers.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "handlers.go",
			Data:       di.Config,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/main.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "main.go",
			Data:       nil,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/server.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "server.go",
			Data:       di.Config.Backend,
		},
		// Deploy files.
		{
			EmbedFile:  "templates/deploy/docker-compose.yml.gotmpl",
			OutputFile: "docker-compose.yml",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  "templates/deploy/Dockerfile.gotmpl",
			OutputFile: "Dockerfile",
			Data:       di.Config,
		},
		{
			EmbedFile:  "templates/deploy/dockerignore.gotmpl",
			OutputFile: ".dockerignore",
			Data:       nil,
		},
		// Misc files.
		{
			EmbedFile:  "templates/misc/gitignore.gotmpl",
			OutputFile: ".gitignore",
			Data:       nil,
		},
	}

	// Check if Air tool is used.
	if di.Config.Tools.IsUseAir {
		// Add Air template to the list.
		templates = append(templates, helpers.EmbedTemplate{
			EmbedFile:  "templates/misc/air.toml.gotmpl",
			OutputFile: ".air.toml",
			Data:       di.Config.Tools,
		})
	} else {
		// Add Makefile to the list.
		templates = append(templates, helpers.EmbedTemplate{
			EmbedFile:  "templates/misc/Makefile.gotmpl",
			OutputFile: "Makefile",
			Data:       di.Config,
		})
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, templates)
}

// createFrontendFiles creates frontend files.
func createFrontendFiles(di *injectors.Injector) error {
	// Check if the frontend is configured to use HTMX.
	if di.Config.Frontend.IsUseHTMX {
		// Generate HTMX files.
		if err := generateHTMXFiles(di); err != nil {
			return err
		}
	}

	// Check which CSS framework is configured for the frontend.
	switch di.Config.Frontend.CSSFramework {
	case "daisyui", "flowbite", "tailwindcss":
		// Generate Tailwind CSS files for supported frameworks.
		if err := generateTailwindCSSFiles(di); err != nil {
			return err
		}
	case "unocss":
		// Generate Uno CSS files.
		if err := generateUnoCSSFiles(di); err != nil {
			return err
		}
	}

	return nil
}

// generateHTMXFiles generates HTMX files.
func generateHTMXFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	generateFiles := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/frontend/htmx/styles.scss.gotmpl",
			OutputFile: "web/src/styles.scss",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/htmx/scripts.js.gotmpl",
			OutputFile: "web/src/scripts.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/htmx/package.json.gotmpl",
			OutputFile: "package.json",
			Data:       di.Config.Frontend,
		},
	}

	// Check if Templ is used.
	if di.Config.Tools.IsUseTempl {
		generateFiles = append(generateFiles,
			helpers.EmbedTemplate{
				EmbedFile:  "templates/frontend/htmx/main.templ.gotmpl",
				OutputFile: "templates/main.templ",
				Data:       di.Config,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/frontend/htmx/index.templ.gotmpl",
				OutputFile: "templates/pages/index.templ",
				Data:       nil,
			},
		)
	} else {
		// Define the files to be copied.
		copyFiles := []helpers.EmbedFile{
			{
				EmbedFile:  "templates/frontend/htmx/main.html",
				OutputFile: "templates/main.html",
			},
			{
				EmbedFile:  "templates/frontend/htmx/index.html",
				OutputFile: "templates/pages/index.html",
			},
		}

		// Copy files from the embed file system.
		if err := helpers.CopyFilesFromEmbedFS(di.Attachments.Templates, copyFiles); err != nil {
			return err
		}
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, generateFiles)
}

// generateTailwindCSSFiles generates Tailwind CSS files.
func generateTailwindCSSFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/frontend/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/tailwind.config.js.gotmpl",
			OutputFile: "tailwind.config.js",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}

// generateUnoCSSFiles generates Uno CSS files.
func generateUnoCSSFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/frontend/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/uno.config.ts.gotmpl",
			OutputFile: "uno.config.ts",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}

// installDependencies installs dependencies.
func installDependencies(di *injectors.Injector) error {
	// Define the commands to be executed.
	commands := []helpers.Command{}

	// Check if Templ is used.
	if di.Config.Tools.IsUseTempl {
		// Add Templ generate command.
		commands = append(commands, helpers.Command{
			Name:       "templ",
			Options:    []string{"generate"},
			SkipOutput: true,
		})
	}

	// Add Go mod tidy command.
	commands = append(commands, helpers.Command{
		Name:       "go",
		Options:    []string{"mod", "tidy"},
		SkipOutput: true,
	})

	// Check if Bun is used.
	if di.Config.Tools.IsUseBun {
		// Add Bun install command.
		commands = append(commands, helpers.Command{
			Name:       "bun",
			Options:    []string{"install"},
			SkipOutput: true,
		})
	} else {
		// Add NPM install command.
		commands = append(commands, helpers.Command{
			Name:       "npm",
			Options:    []string{"install"},
			SkipOutput: true,
		})
	}

	return helpers.Execute(commands)
}