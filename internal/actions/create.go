package actions

import (
	"context"
	"fmt"

	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/gowebly/gowebly/v3/internal/helpers"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
	"github.com/gowebly/gowebly/v3/internal/variables"
)

// CreateProjectAction creates a new project in a goroutine with context support.
func CreateProjectAction(ctx context.Context, cancel context.CancelFunc, di *injectors.Injector, errCh chan<- error) {
	defer cancel()

	// Ordered list of project creation steps.
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

	for _, action := range actions {
		if ctx.Err() != nil {
			errCh <- ctx.Err()
			return
		}

		if err := action.fn(di); err != nil {
			errCh <- fmt.Errorf(messages.ErrorGoroutineActionNotSuccess, action.op, err)
			return
		}
	}

	if ctx.Err() != nil {
		errCh <- ctx.Err()
		return
	}

	errCh <- nil
}

// createProjectFolders creates the standard project folder structure.
func createProjectFolders(di *injectors.Injector) error {
	folders := []string{"assets", "static/images", "templates/pages"}

	return helpers.MakeFolders(folders...)
}

// copyStaticFiles copies embedded static assets to the project.
func copyStaticFiles(di *injectors.Injector) error {
	// Favicon and PWA manifest files.
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
			EmbedFile:  "static/favicon.png",
			OutputFile: "static/apple-touch-icon.png",
		},
		{
			EmbedFile:  "static/favicon.svg",
			OutputFile: "static/favicon.svg",
		},
		{
			EmbedFile:  "static/favicon.svg",
			OutputFile: "static/manifest-touch-icon.svg",
		},
		{
			EmbedFile:  "static/gowebly.svg",
			OutputFile: "static/images/gowebly.svg",
		},
		{
			EmbedFile:  "static/manifest-desktop-screenshot.jpg",
			OutputFile: "static/manifest-desktop-screenshot.jpg",
		},
		{
			EmbedFile:  "static/manifest-mobile-screenshot.jpg",
			OutputFile: "static/manifest-mobile-screenshot.jpg",
		},
		{
			EmbedFile:  "static/manifest.webmanifest",
			OutputFile: "static/manifest.webmanifest",
		},
	}

	return helpers.CopyFilesFromEmbedFS(di.Attachments.Static, files)
}

// createBackendFiles generates all backend source files and configuration.
func createBackendFiles(di *injectors.Injector) error {
	var goFramework, reactivityLibrary, cssFramework string

	// Lookup human-readable names for the selected options.
	if _, ok := variables.ListGoFrameworks[di.Config.Backend.GoFramework]; ok {
		goFramework = variables.ListGoFrameworks[di.Config.Backend.GoFramework][1]
	}

	if _, ok := variables.ListReactivityLibraries[di.Config.Frontend.ReactivityLibrary]; ok {
		reactivityLibrary = variables.ListReactivityLibraries[di.Config.Frontend.ReactivityLibrary][1]
	}

	if _, ok := variables.ListCSSFrameworks[di.Config.Frontend.CSSFramework]; ok {
		cssFramework = variables.ListCSSFrameworks[di.Config.Frontend.CSSFramework][1]
	}

	templates := []helpers.EmbedTemplate{
		// Core backend files
		{
			EmbedFile:  "templates/backend/go.mod.gotmpl",
			OutputFile: "go.mod",
			Data:       di.Config,
		},
		{
			EmbedFile:  "templates/backend/main.go.gotmpl",
			OutputFile: "main.go",
			Data:       nil,
		},
		// Framework-specific handlers and server setup
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/handlers.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "handlers.go",
			Data:       di.Config,
		},
		{
			EmbedFile:  fmt.Sprintf("templates/backend/%s/server.go.gotmpl", di.Config.Backend.GoFramework),
			OutputFile: "server.go",
			Data:       di.Config,
		},
		// Deployment configuration
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
		// Project documentation and config files
		{
			EmbedFile:  "templates/misc/README.md.gotmpl",
			OutputFile: "README.md",
			Data: struct {
				Backend                                      *config.Backend
				Frontend                                     *config.Frontend
				Tools                                        *config.Tools
				GoFramework, ReactivityLibrary, CSSFramework string
			}{
				Backend:           di.Config.Backend,
				Frontend:          di.Config.Frontend,
				Tools:             di.Config.Tools,
				GoFramework:       goFramework,
				ReactivityLibrary: reactivityLibrary,
				CSSFramework:      cssFramework,
			},
		},
		{
			EmbedFile:  "templates/misc/prettier.config.js.gotmpl",
			OutputFile: "prettier.config.js",
			Data:       di.Config.Frontend,
		},
		{
			EmbedFile:  "templates/misc/dockerignore.gotmpl",
			OutputFile: ".dockerignore",
			Data:       nil,
		},
		{
			EmbedFile:  "templates/misc/gitignore.gotmpl",
			OutputFile: ".gitignore",
			Data:       nil,
		},
		{
			EmbedFile:  "templates/misc/prettierignore.gotmpl",
			OutputFile: ".prettierignore",
			Data:       nil,
		},
	}

	if di.Config.Tools.IsUseAir {
		templates = append(templates, helpers.EmbedTemplate{
			EmbedFile:  "templates/misc/air.toml.gotmpl",
			OutputFile: ".air.toml",
			Data:       di.Config,
		})
	} else {
		templates = append(templates, helpers.EmbedTemplate{
			EmbedFile:  "templates/misc/Makefile.gotmpl",
			OutputFile: "Makefile",
			Data:       di.Config,
		})
	}

	if di.Config.Tools.IsUseGolangCILint {
		templates = append(templates, helpers.EmbedTemplate{
			EmbedFile:  "templates/misc/golangci.yml.gotmpl",
			OutputFile: ".golangci.yml",
			Data:       di.Config,
		})
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, templates)
}

// createFrontendFiles generates frontend assets and templates.
func createFrontendFiles(di *injectors.Injector) error {
	if err := generateHTMXFiles(di); err != nil {
		return err
	}

	switch di.Config.Frontend.CSSFramework {
	case "daisyui", "flowbite", "prelineui", "tailwindcss":
		if err := generateTailwindCSSFiles(di); err != nil {
			return err
		}
	case "unocss":
		if err := generateUnoCSSFiles(di); err != nil {
			return err
		}
	}

	return nil
}

// generateHTMXFiles creates HTMX-related frontend files.
func generateHTMXFiles(di *injectors.Injector) error {
	generateFiles := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/frontend/styles.css.gotmpl",
			OutputFile: "assets/styles.css",
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
		copyFiles := []helpers.EmbedFile{
			{
				EmbedFile:  "templates/frontend/main.html.gotmpl",
				OutputFile: "templates/main.html",
			},
			{
				EmbedFile:  "templates/frontend/index.html.gotmpl",
				OutputFile: "templates/pages/index.html",
			},
		}

		if err := helpers.CopyFilesFromEmbedFS(di.Attachments.Templates, copyFiles); err != nil {
			return err
		}
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, generateFiles)
}

// generateTailwindCSSFiles creates Tailwind CSS configuration files.
func generateTailwindCSSFiles(di *injectors.Injector) error {
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/misc/postcssrc.gotmpl",
			OutputFile: ".postcssrc",
			Data:       di.Config.Frontend,
		},
	}

	return helpers.GenerateFilesByTemplateFromEmbedFS(di.Attachments.Templates, files)
}

// generateUnoCSSFiles creates Uno CSS configuration files.
func generateUnoCSSFiles(di *injectors.Injector) error {
	files := []helpers.EmbedTemplate{
		{
			EmbedFile:  "templates/misc/postcssrc.gotmpl",
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

// installDependencies runs all necessary dependency installation commands.
func installDependencies(di *injectors.Injector) error {
	commands := make([]helpers.Command, 0, 4)

	if di.Config.Tools.IsUseTempl {
		commands = append(commands, helpers.Command{
			Name:       "templ",
			Options:    []string{"generate"},
			SkipOutput: true,
		})
	}

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

	frontendRuntimeEnv := "npm"
	if di.Config.Tools.IsUseBun {
		frontendRuntimeEnv = "bun"
	}

	commands = append(commands,
		helpers.Command{
			Name:       frontendRuntimeEnv,
			Options:    []string{"install"},
			SkipOutput: true,
		}, helpers.Command{
			Name:       frontendRuntimeEnv,
			Options:    []string{"run", "fmt"},
			SkipOutput: true,
		}, helpers.Command{
			Name:       frontendRuntimeEnv,
			Options:    []string{"run", "build"},
			SkipOutput: true,
		},
	)

	return helpers.Execute(commands)
}
