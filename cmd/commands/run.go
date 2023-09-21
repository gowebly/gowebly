package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Run runs the 'run' cmd command.
func Run(di *injector.Injector) error {
	helpers.PrintStyled("Successfully run your project in a developer mode!", "success", "margin-top")
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(fmt.Sprintf(
		"Go backend: %s, on port %d", di.Config.Backend.Name, di.Config.Backend.Port),
		"info", "margin-left",
	)
	helpers.PrintStyled("htmx: dev (non-production)", "info", "margin-left")
	helpers.PrintStyled("hyperscript: dev (non-production)", "info", "margin-left")
	if di.Config.Frontend.CSS != nil {
		helpers.PrintStyled(
			fmt.Sprintf("%s: dev (non-production)", di.Config.Frontend.CSS.Framework),
			"info", "margin-left",
		)
	}
	helpers.PrintStyled("Next steps:", "", "margin-top-bottom")
	helpers.PrintStyled(
		fmt.Sprintf("Open http://localhost:%d on a browser to see your project", di.Config.Backend.Port),
		"info", "margin-left",
	)
	helpers.PrintStyled("Go backend logs:", "", "margin-top-bottom")

	return helpers.Execute("go", "run", "./...")
}
