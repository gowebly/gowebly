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
	_ = helpers.RemoveFiles("static/htmx.min.js", "static/hyperscript.min.js", "static/styles.css")

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
					filepath.Join("templates", "misc", "dockerignore.tmpl"),
					".dockerignore",
					di.Config.Frontend,
				},
				{
					filepath.Join("templates", "misc", "Dockerfile.tmpl"),
					"Dockerfile",
					di.Config.Backend,
				},
				{
					filepath.Join("templates", "misc", "docker-compose.yml.tmpl"),
					"docker-compose.yml",
					di.Config.Backend,
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

	// Set the default JavaScript runtime environment.
	frontendRuntime := "npm"

	// Check, if the runtime of the frontend part is switched.
	if di.Config.Frontend.RuntimeEnvironment == "bun" {
		frontendRuntime = "bun"
	}

	// Execute system commands.
	if err := helpers.Execute(
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
				[]string{"run", "build:prod"},
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
				"Successfully build your project for the production!", "success", "margin-top",
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
					"htmx '%s', hyperscript '%s'", di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
				),
				"info", "margin-left-2",
			},
			{
				"Next steps:", "", "margin-top-bottom",
			},
			{
				"Run your project by 'docker-compose up' command on your remote server or local machine",
				"info", "margin-left",
			},
			{
				"You can use an auto-generated 'docker-compose.yml' file on the Portainer platform or manually",
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
