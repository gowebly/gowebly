package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Run runs the 'run' cmd command.
func Run(di *injector.Injector) error {
	// Remove previously generated .env and JS files.
	_ = helpers.RemoveFiles(".env", "static/htmx.min.js", "static/hyperscript.min.js")

	// Create a new folder(s).
	if err := helpers.MakeFolders("static"); err != nil {
		return err
	}

	// Download minified version of the htmx and hyperscript JS files from CND.
	if err := helpers.DownloadFiles(
		[]helpers.Download{
			{
				fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HTMXNameOfCDNRepository, di.Config.Frontend.HTMX,
				),
				"htmx.min.js",
				"static",
			},
			{
				fmt.Sprintf(
					"%s/%s@%s",
					constants.LinkToUnpkgCDN, constants.HyperscriptNameOfCDNRepository, di.Config.Frontend.Hyperscript,
				),
				"hyperscript.min.js",
				"static",
			},
		},
	); err != nil {
		return err
	}

	// Re-create .env file from the template.
	if err := helpers.GenerateFromEmbedFS(
		di.Attachments.Templates,
		[]helpers.EmbedTemplate{
			{
				"templates/misc/env.tmpl",
				".env", "", di.Config.Backend,
			},
		},
	); err != nil {
		return err
	}

	// Start executing commands.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				Name: "go", Options: []string{"mod", "tidy"}, SkipOutput: true,
			},
			{
				Name: "go", Options: []string{"run", "./..."}, SkipOutput: false,
			},
		},
	); err != nil {
		return err
	}

	// Header message.
	helpers.PrintStyled("Successfully run your project in a developer mode!", "success", "margin-top")

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

	if di.Config.Frontend.CSS == nil {
		helpers.PrintStyled("Frontend: default", "info", "margin-left")
	} else {
		helpers.PrintStyled(
			fmt.Sprintf("Frontend: %s ('dev', non-production)", di.Config.Frontend.CSS.Framework),
			"info", "margin-left",
		)
	}

	helpers.PrintStyled(
		"htmx ('dev', non-production), hyperscript ('dev', non-production)",
		"info", "margin-left-2",
	)

	// Next steps message.
	helpers.PrintStyled("Next steps:", "", "margin-top-bottom")
	helpers.PrintStyled(
		fmt.Sprintf("Open http://localhost:%d on a browser to see your project", di.Config.Backend.Port),
		"info", "margin-left",
	)

	// Backend logs message.
	helpers.PrintStyled("Backend logs:", "", "margin-top-bottom")

	return nil
}
