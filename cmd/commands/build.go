package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Build runs the 'build' cmd command.
func Build(di *injector.Injector) error {
	helpers.PrintStyled("Successfully build your project for the production!", "success", "margin-top")
	helpers.PrintStyled("Project configuration:", "", "margin-top-bottom")
	helpers.PrintStyled(fmt.Sprintf(
		"Go backend: %s, on port %d", di.Config.Backend.Name, di.Config.Backend.Port),
		"info", "margin-left",
	)
	helpers.PrintStyled(fmt.Sprintf(
		"htmx: %s", di.Config.Frontend.HTMX),
		"info", "margin-left",
	)
	helpers.PrintStyled(fmt.Sprintf(
		"hyperscript: %s", di.Config.Frontend.Hyperscript),
		"info", "margin-left",
	)
	if di.Config.Frontend.CSS != nil {
		helpers.PrintStyled(fmt.Sprintf(
			"%s: %s", di.Config.Frontend.CSS.Framework, di.Config.Frontend.CSS.Version),
			"info", "margin-left",
		)
	}
	helpers.PrintStyled("Next steps:", "", "margin-top-bottom")
	helpers.PrintStyled(
		"Run your project by 'docker-compose up -d' command on your remote server or local machine",
		"info", "margin-left",
	)
	helpers.PrintStyled(
		"You can use an auto-generated 'docker-compose.yml' file on the Portainer platform or manually",
		"info", "margin-left",
	)
	helpers.PrintStyled(
		fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
		"warning", "margin-top-bottom",
	)

	return nil
}
