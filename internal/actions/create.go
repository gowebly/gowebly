package actions

import (
	"context"
	"fmt"

	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
)

// CreateProjectAction creates a new project.
func CreateProjectAction(ctx context.Context, di *injectors.Injector, errCh chan<- error) {
	// Create project folders.
	if err := createProjectFolders(); err != nil {
		errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, "create project folders", err)
		return
	}

	// Create backend files.
	if err := createBackendFiles(di); err != nil {
		errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, "generate backend files", err)
		return
	}

	// Create frontend files.
	if err := createFrontendFiles(di); err != nil {
		errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, "generate frontend files", err)
		return
	}

	select {
	case <-ctx.Done():
		errCh <- ctx.Err()
	default:
		errCh <- nil
	}
}

// createProjectFolders creates project folders.
func createProjectFolders() error {
	// Define a slice of folder paths.
	folders := []string{
		"web/static",          // folder for static files
		"web/templates/pages", // folder for page templates
		"web/src/scss",        // folder for SCSS files
		"web/src/js",          // folder for JavaScript files
	}

	// Call the MakeFolders function from the helpers package.
	// Pass the folders slice as variadic arguments.
	return helpers.MakeFolders(folders...)
}

// createBackendFiles creates backend files.
func createBackendFiles(di *injectors.Injector) error {
	// Define embed templates.
	templates := []helpers.EmbedTemplate{
		// Backend files.
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/go.mod.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "go.mod",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/server.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "server.go",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/handlers.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "handlers.go",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/main.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "main.go",
			Data:       nil,
		},
		// Deploy files.
		{
			EmbedFile:  "templates/deploy/dockerignore.gotmpl",
			OutputFile: ".dockerignore",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  "templates/deploy/Dockerfile.gotmpl",
			OutputFile: "Dockerfile",
			Data:       di.Config.Backend,
		},
		{
			EmbedFile:  "templates/deploy/docker-compose.yml.gotmpl",
			OutputFile: "docker-compose.yml",
			Data:       di.Config.Backend,
		},
		// Misc files.
		{
			EmbedFile:  "templates/misc/gitignore.gotmpl",
			OutputFile: ".gitignore",
			Data:       nil,
		},
		{
			EmbedFile:  "templates/misc/air.toml.gotmpl",
			OutputFile: "air.toml",
			Data:       nil,
		},
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
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/htmx/styles.scss.gotmpl",
			OutputFile: "web/src/scss/styles.scss",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/htmx/scripts.js.gotmpl",
			OutputFile: "web/src/js/scripts.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/htmx/package.json.gotmpl",
			OutputFile: "package.json",
			Data:       di.Config.Frontend,
		},
	}

	// Check if Templ is used.
	if di.Config.Tools.IsUseTempl {
		files = append(files,
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/main.templ.gotmpl",
				OutputFile: "web/templates/main.templ",
				Data:       di.Config.Frontend,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/index.templ.gotmpl",
				OutputFile: "web/templates/pages/index.templ",
				Data:       di.Config.Frontend,
			},
		)
	} else {
		files = append(files,
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/main.html.gotmpl",
				OutputFile: "web/templates/main.html",
				Data:       di.Config.Frontend,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/htmx/index.html.gotmpl",
				OutputFile: "web/templates/pages/index.html",
				Data:       di.Config.Frontend,
			},
		)
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}

// generateTailwindCSSFiles generates Tailwind CSS files.
func generateTailwindCSSFiles(di *injectors.Injector) error {
	// Define the files to be generated.
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/css/tailwind.config.js.gotmpl",
			OutputFile: "tailwind.config.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/css/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
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
			EmbedFile:  "templates/css/uno.config.ts.gotmpl",
			OutputFile: "uno.config.ts",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/css/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}
