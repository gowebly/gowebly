package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// GoFrameworkSelect runs the Go framework select.
func GoFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	// Define options for the select.
	options := make([]huh.Option[string], 0, len(variables.ListGoFrameworks))

	// Add options to the select.
	for _, framework := range variables.ListGoFrameworks {
		options = append(options, huh.NewOption(framework[1], framework[0]))
	}

	return huh.NewSelect[string]().
		Title(messages.FormGoFrameworkTitle).
		Description(messages.FormGoFrameworkDescription).
		Options(options...).
		Value(&di.Config.Backend.GoFramework)
}

// ReactivityLibrarySelect runs the reactivity library select.
func ReactivityLibrarySelect(di *injectors.Injector) *huh.Select[string] {
	// Define options for the select.
	options := make([]huh.Option[string], 0, len(variables.ListReactivityLibraries))

	// Add options to the select.
	for _, library := range variables.ListReactivityLibraries {
		options = append(options, huh.NewOption(library[1], library[0]))
	}

	return huh.NewSelect[string]().
		Title(messages.FormReactivityLibraryTitle).
		Description(messages.FormReactivityLibraryDescription).
		Options(options...).
		Value(&di.Config.Frontend.ReactivityLibrary)
}

// CSSFrameworkSelect runs the CSS framework select.
func CSSFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	// Define options for the select.
	options := make([]huh.Option[string], 0, len(variables.ListCSSFrameworks))

	// Add options to the select.
	for _, framework := range variables.ListCSSFrameworks {
		options = append(options, huh.NewOption(framework[1], framework[0]))
	}

	return huh.NewSelect[string]().
		Title(messages.FormCSSFrameworkTitle).
		Description(messages.FormCSSFrameworkDescription).
		Options(options...).
		Value(&di.Config.Frontend.CSSFramework)
}
