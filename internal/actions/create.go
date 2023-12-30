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

	// Copy static files from embed FS.
	if err := copyStaticFiles(di); err != nil {
		errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, "copy static files", err)
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

	return helpers.MakeFolders(folders...)
}

// copyStaticFiles copies static files.
func copyStaticFiles(di *injectors.Injector) error {
	// Define the files to be copied.
	files := []helpers.EmbedFile{
		{
			EmbedFile:  "static/favicon.ico",
			OutputFile: "web/static/favicon.ico",
		},
		{
			EmbedFile:  "static/favicon.png",
			OutputFile: "web/static/favicon.png",
		},
		{
			EmbedFile:  "static/favicon.svg",
			OutputFile: "web/static/favicon.svg",
		},
		{
			EmbedFile:  "static/logo.svg",
			OutputFile: "web/static/logo.svg",
		},
		{
			EmbedFile:  "static/manifest-desktop-screenshot.jpeg",
			OutputFile: "web/static/manifest-desktop-screenshot.jpeg",
		},
		{
			EmbedFile:  "static/manifest-mobile-screenshot.jpeg",
			OutputFile: "web/static/manifest-mobile-screenshot.jpeg",
		},
		{
			EmbedFile:  "static/manifest.json",
			OutputFile: "web/static/manifest.json",
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
			EmbedFile:  "templates/misc/air.toml.gotmpl",
			OutputFile: "air.toml",
			Data:       di.Config.Tools,
		},
		{
			EmbedFile:  "templates/misc/gitignore.gotmpl",
			OutputFile: ".gitignore",
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
	generateFiles := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/frontend/htmx/styles.scss.gotmpl",
			OutputFile: "web/src/scss/styles.scss",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/htmx/scripts.js.gotmpl",
			OutputFile: "web/src/js/scripts.js",
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
				OutputFile: "web/templates/main.templ",
				Data:       di.Config,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/frontend/htmx/index.templ.gotmpl",
				OutputFile: "web/templates/pages/index.templ",
				Data:       nil,
			},
		)
	} else {
		// Define the files to be copied.
		copyFiles := []helpers.EmbedFile{
			{
				EmbedFile:  "templates/frontend/htmx/main.html",
				OutputFile: "web/templates/main.html",
			},
			{
				EmbedFile:  "templates/frontend/htmx/index.html",
				OutputFile: "web/templates/pages/index.html",
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
			EmbedFile:  "templates/frontend/css/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/css/tailwind.config.js.gotmpl",
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
			EmbedFile:  "templates/frontend/css/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/css/uno.config.ts.gotmpl",
			OutputFile: "uno.config.ts",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}
