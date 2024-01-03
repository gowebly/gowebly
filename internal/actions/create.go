package actions

import (
	"context"
	"fmt"

	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
)

// CreateProjectAction creates a new project.
func CreateProjectAction(ctx context.Context, cancel context.CancelFunc, di *injectors.Injector, errCh chan<- error) {
	// Define the list of actions to execute.
	actions := []struct {
		op string
		fn func(*injectors.Injector) error
	}{
		{"create project folders", createProjectFolders},
		{"copy static files", copyStaticFiles},
		{"generate backend files", createBackendFiles},
		{"generate frontend files", createFrontendFiles},
		{"install project dependencies", installDependencies},
	}

	// Execute each action in the list
	for _, action := range actions {
		// Check context.
		if ctx.Err() != nil {
			cancel()
			errCh <- ctx.Err()
			return
		}

		// Run action.
		if err := action.fn(di); err != nil {
			cancel()
			errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, action.op, err)
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
	folders := []string{"assets", "static", "templates/pages"}

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
			Data:       di.Config,
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
	// Generate HTMX files.
	if err := generateHTMXFiles(di); err != nil {
		return err
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
			EmbedFile:  "templates/frontend/styles.scss.gotmpl",
			OutputFile: "assets/styles.scss",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/scripts.js.gotmpl",
			OutputFile: "assets/scripts.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/frontend/package.json.gotmpl",
			OutputFile: "package.json",
			Data:       di.Config.Frontend,
		},
	}

	// Check if Templ is used.
	if di.Config.Tools.IsUseTempl {
		generateFiles = append(generateFiles,
			helpers.EmbedTemplate{
				EmbedFile:  "templates/frontend/main.templ.gotmpl",
				OutputFile: "templates/main.templ",
				Data:       di.Config,
			},
			helpers.EmbedTemplate{
				EmbedFile:  "templates/frontend/index.templ.gotmpl",
				OutputFile: "templates/pages/index.templ",
				Data:       nil,
			},
		)
	} else {
		// Define the files to be copied.
		copyFiles := []helpers.EmbedFile{
			{
				EmbedFile:  "templates/frontend/main.html",
				OutputFile: "templates/main.html",
			},
			{
				EmbedFile:  "templates/frontend/index.html",
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
	commands := make([]helpers.Command, 0, 4)

	// Check if Templ is used.
	if di.Config.Tools.IsUseTempl {
		// Generate templates.
		commands = append(commands, helpers.Command{
			Name:       "templ",
			Options:    []string{"generate"},
			SkipOutput: true,
		})
	}

	// Install Go dependencies.
	commands = append(commands,
		helpers.Command{
			Name:       "go",
			Options:    []string{"mod", "tidy"},
			SkipOutput: true,
		},
		helpers.Command{
			Name:       "go",
			Options:    []string{"fmt"},
			SkipOutput: true,
		},
	)

	// Define the frontend runtime environment.
	frontendRuntimeEnv := "npm"

	// Check if Bun is used.
	if di.Config.Tools.IsUseBun {
		// Set the frontend runtime environment to Bun.
		frontendRuntimeEnv = "bun"
	}

	// Install frontend dependencies.
	commands = append(commands,
		helpers.Command{
			Name:       frontendRuntimeEnv,
			Options:    []string{"install"},
			SkipOutput: true,
		}, helpers.Command{
			Name:       frontendRuntimeEnv,
			Options:    []string{"run", "build"},
			SkipOutput: true,
		},
	)

	return helpers.Execute(commands)
}
