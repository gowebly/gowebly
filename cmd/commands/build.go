package commands

import (
	"fmt"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Build runs the 'build' cmd command.
func Build(di *injector.Injector) error {
	// Header message.
	helpers.PrintStyled(
		"Downloading and preparing a minified versions of the frontend part... Please wait!",
		"info", "margin-top",
	)

	// Remove previously generated .env and JS files.
	_ = helpers.RemoveFiles(".env", "static/htmx.min.js", "static/hyperscript.min.js", "static/styles.css")

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

	// Re-create .env file from the template.
	if err := helpers.GenerateFilesByTemplateFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedTemplate{
			{
				filepath.Join("templates", "misc", "env.tmpl"),
				".env",
				di.Config.Backend,
			},
		},
	); err != nil {
		return err
	}

	// Execute system commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				true, "npm", []string{"run", "build:prod"},
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
