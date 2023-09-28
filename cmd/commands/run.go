package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// Run runs the 'run' cmd command.
func Run(di *injectors.Injector) error {
	// Header message.
	helpers.PrintStyled(
		"Re-generate styles of the frontend part... Please wait!",
		"wait", "margin-top",
	)

	// Remove previously generated .env and JS files.
	_ = helpers.RemoveFiles("static/styles.css")

	// Re-generate styles of the frontend part.
	if err := helpers.Execute(
		[]helpers.Command{
			{
				true, "npm", []string{"run", "build:dev"}, nil,
			},
		},
	); err != nil {
		return err
	}

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				"Successfully run your project in a developer mode!",
				"success", "margin-top",
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
				fmt.Sprintf("Frontend: %s 'dev' (non-production)", di.Config.Frontend.CSSFramework),
				"info", "margin-left",
			},
			{
				fmt.Sprintf(
					"htmx '%s', hyperscript '%s'",
					di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
				),
				"info", "margin-left-2",
			},
			{
				"Next steps:", "", "margin-top-bottom",
			},
			{
				fmt.Sprintf(
					"Open http://localhost:%d on a browser to see your project",
					di.Config.Backend.Port,
				),
				"info", "margin-left",
			},
			{
				"Backend logs:", "", "margin-top-bottom",
			},
		},
	)

	return helpers.Execute(
		[]helpers.Command{
			{
				false,
				"go",
				[]string{"run", "./..."},
				[]string{
					fmt.Sprintf("BACKEND_PORT=%d", di.Config.Backend.Port),
					fmt.Sprintf("BACKEND_READ_TIMEOUT=%d", di.Config.Backend.Timeout.Read),
					fmt.Sprintf("BACKEND_WRITE_TIMEOUT=%d", di.Config.Backend.Timeout.Write),
				},
			},
		},
	)
}
