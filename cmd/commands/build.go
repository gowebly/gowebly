package commands

import (
	"fmt"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// Build runs the 'build' cmd command.
func Build(di *injectors.Injector, flag string) error {
	// Remove previously generated files.
	_ = helpers.RemoveFiles("static/htmx.min.js", "static/hyperscript.min.js", "static/styles.css", "static/scripts.js")

	// Check, if the second flag is set.
	if flag != "" {
		// Skip Docker part message.
		helpers.PrintStyled(
			"Re-generation process of the Docker files was skipped (by the '--skip-docker' flag)!",
			"wait", "margin-top",
		)
	} else {
		// Docker part message.
		helpers.PrintStyled(
			"Preparing Docker files for the deploy part... Please wait!",
			"wait", "margin-top",
		)

		// Remove previously generated files.
		_ = helpers.RemoveFiles(".dockerignore", "Dockerfile", "docker-compose.yml")

		// Create Docker part files from templates.
		if err := helpers.GenerateFilesByTemplateFromEmbedFS(
			di.Attachments.Templates,
			[]helpers.EmbedTemplate{
				{
					EmbedFile:  "templates/misc/dockerignore.tmpl",
					OutputFile: ".dockerignore",
					Data:       di.Config.Frontend,
				},
				{
					EmbedFile:  "templates/misc/Dockerfile.tmpl",
					OutputFile: "Dockerfile",
					Data:       di.Config.Backend,
				},
				{
					EmbedFile:  "templates/misc/docker-compose.yml.tmpl",
					OutputFile: "docker-compose.yml",
					Data:       di.Config.Backend,
				},
			},
		); err != nil {
			return err
		}
	}

	// Header message.
	helpers.PrintStyled(
		"Downloading and preparing a minified versions of the frontend part... Please wait!",
		"wait", "",
	)

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

	// Set the default JavaScript runtime environment.
	frontendRuntime := "npm"

	// Check, if the runtime of the frontend part is switched.
	if di.Config.Frontend.RuntimeEnvironment == "bun" {
		frontendRuntime = "bun"
	}

	// Execute system commands.
	switch di.Config.Backend.TemplateEngine {
	case "templ":
		// Build project with the a-h/templ template engine.
		// See https://github.com/a-h/templ for more information.
		if err := helpers.ExecuteInGoroutine(
			[]helpers.Command{
				{
					Name:       "go",
					Options:    []string{"mod", "tidy"},
					SkipOutput: true,
					EnvVars:    nil,
				},
				{
					Name:       "templ",
					Options:    []string{"generate"},
					SkipOutput: true,
					EnvVars:    nil,
				},
				{
					Name:       frontendRuntime,
					Options:    []string{"run", "build:prod"},
					SkipOutput: true,
					EnvVars:    nil,
				},
			},
		); err != nil {
			return err
		}
	default:
		// Build project with the default template/html (built-in) template engine.
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
					Options:    []string{"run", "build:prod"},
					SkipOutput: true,
					EnvVars:    nil,
				},
			},
		); err != nil {
			return err
		}
	}

	// Execute system commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				Name:       "go",
				Options:    []string{"fmt"},
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
				Text: "Successfully build your project for the production!", State: "success", Style: "margin-top",
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
				Text:  fmt.Sprintf("Template engine: %s", di.Config.Backend.TemplateEngine),
				State: "info", Style: "margin-left-2",
			},
			{
				Text:  fmt.Sprintf("Frontend: %s", di.Config.Frontend.CSSFramework),
				State: "info", Style: "margin-left",
			},
			{
				Text: fmt.Sprintf(
					"htmx '%s', hyperscript '%s'", di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
				),
				State: "info", Style: "margin-left-2",
			},
			{
				Text: "Next steps:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  "Run your project by 'docker-compose up' command on your remote server or local machine",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "You can use an auto-generated 'docker-compose.yml' file on the Portainer platform or manually",
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
