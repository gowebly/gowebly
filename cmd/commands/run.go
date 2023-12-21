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

	// Remove previously generated style files.
	_ = helpers.RemoveFiles("static/styles.css", "static/scripts.js")

	// Print block of messages.
	helpers.PrintStyledBlock(
		[]helpers.StyledOutput{
			{
				Text:  "Successfully run your project in a developer mode!",
				State: "success", Style: "margin-top",
			},
			{
				Text: "Project configuration:", State: "", Style: "margin-top-bottom",
			},
			{
				Text:  fmt.Sprintf("Backend: %s", di.Config.Backend.GoFramework),
				State: "info", Style: "margin-left",
			},
			{
				Text: fmt.Sprintf(
					"Server port is %d, timeout (in seconds): %d for read, %d for write",
					di.Config.Backend.Port, di.Config.Backend.Timeout.Read, di.Config.Backend.Timeout.Write,
				),
				State: "info", Style: "margin-left-2",
			},
			{
				Text:  fmt.Sprintf("Template engine: %s", di.Config.Backend.TemplateEngine),
				State: "info", Style: "margin-left-2",
			},
			{
				Text:  fmt.Sprintf("Frontend: %s 'dev' (non-production)", di.Config.Frontend.CSSFramework),
				State: "info", Style: "margin-left",
			},
			{
				Text: fmt.Sprintf(
					"htmx '%s', hyperscript '%s'",
					di.Config.Frontend.HTMX, di.Config.Frontend.Hyperscript,
				),
				State: "info", Style: "margin-left-2",
			},
			{
				Text: "Next steps:", State: "", Style: "margin-top-bottom",
			},
			{
				Text: fmt.Sprintf(
					"Open http://localhost:%d on a browser to see your project",
					di.Config.Backend.Port,
				),
				State: "info", Style: "margin-left",
			},
			{
				Text: "Backend logs:", State: "", Style: "margin-top-bottom",
			},
		},
	)

	// Set the default JavaScript runtime environment.
	frontendRuntime := "npm"

	// Check, if the runtime of the frontend part is switched.
	if di.Config.Frontend.RuntimeEnvironment == "bun" {
		frontendRuntime = "bun"
	}

	// Execute system commands.
	switch di.Config.Backend.TemplateEngine {
	case "templ":
		// Execute system commands with a-h/templ template engine.
		// See https://github.com/a-h/templ for more information.
		return helpers.ExecuteInGoroutine(
			[]helpers.Command{
				{
					Name:       frontendRuntime,
					Options:    []string{"run", "watch"},
					SkipOutput: true,
				},
				{
					Name:       "templ",
					Options:    []string{"generate", "--watch"},
					SkipOutput: true,
				},
				{
					Name:    "go",
					Options: []string{"run", "."},
					EnvVars: []string{
						fmt.Sprintf("BACKEND_PORT=%d", di.Config.Backend.Port),
						fmt.Sprintf("BACKEND_READ_TIMEOUT=%d", di.Config.Backend.Timeout.Read),
						fmt.Sprintf("BACKEND_WRITE_TIMEOUT=%d", di.Config.Backend.Timeout.Write),
					},
				},
			},
		)
	default:
		// Execute system commands with the default template/html (built-in) template engine.
		return helpers.ExecuteInGoroutine(
			[]helpers.Command{
				{
					Name:       frontendRuntime,
					Options:    []string{"run", "watch"},
					SkipOutput: true,
				},
				{
					Name:    "go",
					Options: []string{"run", "."},
					EnvVars: []string{
						fmt.Sprintf("BACKEND_PORT=%d", di.Config.Backend.Port),
						fmt.Sprintf("BACKEND_READ_TIMEOUT=%d", di.Config.Backend.Timeout.Read),
						fmt.Sprintf("BACKEND_WRITE_TIMEOUT=%d", di.Config.Backend.Timeout.Write),
					},
				},
			},
		)
	}
}
