package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injector"
)

// Build runs the 'build' cmd command.
func Build(di *injector.Injector) error {
	// Header message.
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

	if di.Config.Frontend.CSS == nil {
		helpers.PrintStyled("Frontend: default", "info", "margin-left")
	} else {
		helpers.PrintStyled(
			fmt.Sprintf(
				"Frontend: %s ('%s')",
				di.Config.Frontend.CSS.Framework, di.Config.Frontend.CSS.Version,
			),
			"info", "margin-left",
		)
	}

	helpers.PrintStyled(
		fmt.Sprintf("htmx ('%s'), hyperscript ('%s')", di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript),
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
