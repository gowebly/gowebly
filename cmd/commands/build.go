package commands

import (
	"fmt"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Build runs the 'build' cmd command.
func Build(di *injector.Injector, flags []string) error {
	// Remove previously generated files.
	_ = helpers.RemoveFiles("static/htmx.min.js", "static/hyperscript.min.js", "static/styles.css")

	// Define the Docker part marker.
	var skipDockerPart bool

	// Check, if flag set more than 1.
	if len(flags) > 1 {
		// Check, if the second flag is '--skip-docker'.
		if flags[1] == "--skip-docker" {
			// Set the Docker part marker to 'true'.
			skipDockerPart = true

			// Header message.
			helpers.PrintStyled(
				"Re-generation process of the 'Dockerfile' and 'docker-compose.yml' was skipped (by the '--skip-docker' flag)!",
				"info", "margin-top",
			)
		}
	}

	// Check, if user want to skip re-generation process of the 'Dockerfile' and
	// 'docker-compose.yml'.
	if !skipDockerPart {
		// Header message.
		helpers.PrintStyled(
			"Preparing Docker files ('Dockerfile' and 'docker-compose.yml') for the deploy part... Please wait!",
			"info", "margin-top",
		)

		// Remove previously generated files.
		_ = helpers.RemoveFiles("Dockerfile", "docker-compose.yml")

		// Create Docker part files from templates.
		if err := helpers.GenerateFilesByTemplateFromEmbedFS(
			di.Attachments.Templates,
			[]helpers.EmbedTemplate{
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
		"info", "margin-top",
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

	// Execute system commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				true, "npm", []string{"run", "build:prod"}, nil,
			},
		},
	); err != nil {
		return err
	}

	// Success message.
	helpers.PrintStyled("Successfully build your project for the production!", "success", "margin-top")

	// Project config message.
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(
		fmt.Sprintf("Backend: %s", di.Config.Backend.Name),
		"info", "margin-left",
	)
	helpers.PrintStyled(
		fmt.Sprintf(
			"Server port is %d, timeout (in seconds): %d for read, %d for write",
			di.Config.Backend.Port, di.Config.Backend.Timeout.Read, di.Config.Backend.Timeout.Write,
		),
		"info", "margin-left-2",
	)
	helpers.PrintStyled(
		fmt.Sprintf("Frontend: %s", di.Config.Frontend.CSSFramework),
		"info", "margin-left",
	)
	helpers.PrintStyled(
		fmt.Sprintf("htmx '%s', hyperscript '%s'", di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript),
		"info", "margin-left-2",
	)

	// Next steps message.
	helpers.PrintStyled("Next steps:", "", "margin-top-bottom")
	helpers.PrintStyled(
		"Run your project by 'docker-compose up -d' command on your remote server or local machine",
		"info", "margin-left",
	)
	helpers.PrintStyled(
		"You can use an auto-generated 'docker-compose.yml' file on the Portainer platform or manually",
		"info", "margin-left",
	)

	// Footer message.
	helpers.PrintStyled(
		fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
		"warning", "margin-top-bottom",
	)

	return nil
}
