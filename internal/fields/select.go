package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// GoFrameworkSelect runs the Go framework select.
func GoFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormGoFrameworkTitle).
		Description(messages.FormGoFrameworkDescription).
		Options(
			huh.NewOption(variables.ListGoFrameworks["default"][1], variables.ListGoFrameworks["default"][0]),
			huh.NewOption(variables.ListGoFrameworks["fiber"][1], variables.ListGoFrameworks["fiber"][0]),
			huh.NewOption(variables.ListGoFrameworks["gin"][1], variables.ListGoFrameworks["gin"][0]),
			huh.NewOption(variables.ListGoFrameworks["echo"][1], variables.ListGoFrameworks["echo"][0]),
			huh.NewOption(variables.ListGoFrameworks["chi"][1], variables.ListGoFrameworks["chi"][0]),
			huh.NewOption(variables.ListGoFrameworks["httprouter"][1], variables.ListGoFrameworks["httprouter"][0]),
		).
		Value(&di.Config.Backend.GoFramework)
}

// ReactivityLibrarySelect runs the reactivity library select.
func ReactivityLibrarySelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormReactivityLibraryTitle).
		Description(messages.FormReactivityLibraryDescription).
		Options(
			huh.NewOption(variables.ListReactivityLibraries["htmx"][1], variables.ListReactivityLibraries["htmx"][0]),
			huh.NewOption(variables.ListReactivityLibraries["htmx_hyperscript"][1], variables.ListReactivityLibraries["htmx_hyperscript"][0]),
		).
		Value(&di.Config.Frontend.ReactivityLibrary)
}

// CSSFrameworkSelect runs the CSS framework select.
func CSSFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormCSSFrameworkTitle).
		Description(messages.FormCSSFrameworkDescription).
		Options(
			huh.NewOption(variables.ListCSSFrameworks["default"][1], variables.ListCSSFrameworks["default"][0]),
			huh.NewOption(variables.ListCSSFrameworks["tailwindcss"][1], variables.ListCSSFrameworks["tailwindcss"][0]),
			huh.NewOption(variables.ListCSSFrameworks["daisyui"][1], variables.ListCSSFrameworks["daisyui"][0]),
			huh.NewOption(variables.ListCSSFrameworks["flowbite"][1], variables.ListCSSFrameworks["flowbite"][0]),
			huh.NewOption(variables.ListCSSFrameworks["bootstrap"][1], variables.ListCSSFrameworks["bootstrap"][0]),
			huh.NewOption(variables.ListCSSFrameworks["bulma"][1], variables.ListCSSFrameworks["bulma"][0]),
		).
		Value(&di.Config.Frontend.CSSFramework)
}
