package fields

import (
	"sort"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// sortMapByKey sorts the map by the keys.
func sortMapByKey(inputs map[string][]string) []huh.Option[string] {
	// Define options for the select.
	options := make([]huh.Option[string], 0, len(inputs))

	// Add options to the select.
	for _, input := range inputs {
		options = append(options, huh.NewOption(input[1], input[0]))
	}

	// Sort the options slice by the keys.
	sort.Slice(options, func(i, j int) bool {
		return options[i].Key < options[j].Key
	})

	return options
}

// GoFrameworkSelect runs the Go framework select.
func GoFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormGoFrameworkTitle).
		Description(messages.FormGoFrameworkDescription).
		Options(sortMapByKey(variables.ListGoFrameworks)...).
		Value(&di.Config.Backend.GoFramework)
}

// ReactivityLibrarySelect runs the reactivity library select.
func ReactivityLibrarySelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormReactivityLibraryTitle).
		Description(messages.FormReactivityLibraryDescription).
		Options(sortMapByKey(variables.ListReactivityLibraries)...).
		Value(&di.Config.Frontend.ReactivityLibrary)
}

// CSSFrameworkSelect runs the CSS framework select.
func CSSFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormCSSFrameworkTitle).
		Description(messages.FormCSSFrameworkDescription).
		Options(sortMapByKey(variables.ListCSSFrameworks)...).
		Value(&di.Config.Frontend.CSSFramework)
}
