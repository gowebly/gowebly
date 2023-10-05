package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// Doctor runs the 'doctor' cmd command.
func Doctor(di *injectors.Injector) error {
	// Header message.
	helpers.PrintStyled(
		"Generating a report about your system... Please wait!",
		"wait", "margin-top",
	)

	// Set the default JavaScript runtime environment.
	frontendRuntime := "node"

	// Check, if the runtime of the frontend part is switched.
	if di.Config.Frontend.RuntimeEnvironment == "bun" {
		frontendRuntime = "bun"
	}

	//
	tools := helpers.CheckTools(
		[]helpers.Tool{
			{
				"go", "version",
			},
			{
				"docker", "-v",
			},
			{
				"docker-compose", "-v",
			},
			{
				frontendRuntime, "-v",
			},
		},
	)

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				"Configuration of your system:", "", "margin-top-bottom",
			},
			{
				fmt.Sprintf("gowebly version %s", constants.CLIVersion),
				"success", "margin-left",
			},
			{
				tools["go"][0],
				tools["go"][1], "margin-left",
			},
			{
				tools["docker"][0],
				tools["docker"][1], "margin-left",
			},
			{
				tools["docker-compose"][0],
				tools["docker-compose"][1], "margin-left",
			},
			{
				fmt.Sprintf("%s %s", frontendRuntime, tools[frontendRuntime][0]),
				tools[frontendRuntime][1], "margin-left",
			},
			{
				"Next steps:", "", "margin-top-bottom",
			},
			{
				"If some tools from this list haven't been installed, we strongly recommend installing them manually",
				"info", "margin-left",
			},
			{
				"You can add this information to your issue on GitHub and developers will be able to help you more quickly",
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
