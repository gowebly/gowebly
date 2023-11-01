package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
	"github.com/gowebly/gowebly/internal/injectors"
)

// version represents the current gowebly CLI version.
var version string = "dev"

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

	// Create a list of tools to check.
	toolsList := []helpers.Tool{
		{
			Name: "go", VersionCommand: "version",
		},
		{
			Name: "docker", VersionCommand: "-v",
		},
		{
			Name: "docker-compose", VersionCommand: "-v",
		},
		{
			Name: frontendRuntime, VersionCommand: "-v",
		},
	}

	// Check, if the template engine is templ.
	if di.Config.Backend.TemplateEngine == "templ" {
		// Add templ tool.
		toolsList = append(toolsList, helpers.Tool{
			Name: "templ", VersionCommand: "version",
		})
	}

	// Check tools.
	tools := helpers.CheckTools(toolsList)

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				Text: "Configuration of your system:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  fmt.Sprintf("gowebly version %s", version),
				State: "success", Style: "margin-left",
			},
			{
				Text: tools[0].Output, State: tools[0].Status, Style: "margin-left",
			},
			{
				Text: tools[1].Output, State: tools[1].Status, Style: "margin-left",
			},
			{
				Text: tools[2].Output, State: tools[2].Status, Style: "margin-left",
			},
			{
				Text:  fmt.Sprintf("%s %s", frontendRuntime, tools[3].Output),
				State: tools[3].Status, Style: "margin-left",
			},
		},
	)

	// Check, if the template engine is templ.
	if di.Config.Backend.TemplateEngine == "templ" {
		// Print block of message for templ.
		helpers.PrintStyledBlock(
			[]helpers.StyledOutput{
				{
					Text:  fmt.Sprintf("%s %s", "templ", tools[4].Output),
					State: tools[4].Status, Style: "margin-left",
				},
			},
		)
	}

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				Text: "Next steps:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  "If some tools from this list haven't been installed, we strongly recommend installing them manually",
				State: "info", Style: "margin-left",
			},
			{
				Text:  "You can add this information to your issue on GitHub and developers will be able to help you more quickly",
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
