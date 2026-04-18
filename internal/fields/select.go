package fields

import (
	"sort"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
	"github.com/gowebly/gowebly/v3/internal/variables"
)

func sortMapByKey(inputs map[string][]string) []huh.Option[string] {
	options := make([]huh.Option[string], 0, len(inputs))

	for _, input := range inputs {
		options = append(options, huh.NewOption(input[1], input[0]))
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].Key < options[j].Key
	})

	return options
}

// GoFrameworkSelect prompts for the Go web framework or router selection.
func GoFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormGoFrameworkTitle).
		Description(messages.FormGoFrameworkDescription).
		Options(sortMapByKey(variables.ListGoFrameworks)...).
		Value(&di.Config.Backend.GoFramework)
}

// ReactivityLibrarySelect prompts for the frontend reactivity library (htmx, Alpine.js, etc.).
func ReactivityLibrarySelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormReactivityLibraryTitle).
		Description(messages.FormReactivityLibraryDescription).
		Options(sortMapByKey(variables.ListReactivityLibraries)...).
		Value(&di.Config.Frontend.ReactivityLibrary)
}

// CSSFrameworkSelect prompts for the CSS framework (Tailwind, Bootstrap, etc.).
func CSSFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormCSSFrameworkTitle).
		Description(messages.FormCSSFrameworkDescription).
		Options(sortMapByKey(variables.ListCSSFrameworks)...).
		Value(&di.Config.Frontend.CSSFramework)
}
