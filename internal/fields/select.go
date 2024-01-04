package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
)

func GoFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormGoFrameworkTitle).
		Description(messages.FormGoFrameworkDescription).
		Options(
			huh.NewOption(messages.FormGoFrameworkDefaultKey, messages.FormGoFrameworkDefaultValue),
			huh.NewOption(messages.FormGoFrameworkFiberKey, messages.FormGoFrameworkFiberValue),
			huh.NewOption(messages.FormGoFrameworkGinKey, messages.FormGoFrameworkGinValue),
			huh.NewOption(messages.FormGoFrameworkEchoKey, messages.FormGoFrameworkEchoValue),
			huh.NewOption(messages.FormGoFrameworkHttpRouterKey, messages.FormGoFrameworkHttpRouterValue),
		).
		Value(&di.Config.Backend.GoFramework)
}

func CSSFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormCSSFrameworkTitle).
		Description(messages.FormCSSFrameworkDescription).
		Options(
			huh.NewOption(messages.FormCSSFrameworkDefaultKey, messages.FormCSSFrameworkDefaultValue),
			huh.NewOption(messages.FormCSSFrameworkTailwindCSSKey, messages.FormCSSFrameworkTailwindCSSValue),
			huh.NewOption(messages.FormCSSFrameworkDaisyUIKey, messages.FormCSSFrameworkDaisyUIValue),
			huh.NewOption(messages.FormCSSFrameworkFlowbiteKey, messages.FormCSSFrameworkFlowbiteValue),
			huh.NewOption(messages.FormCSSFrameworkUnoCSSKey, messages.FormCSSFrameworkUnoCSSValue),
			huh.NewOption(messages.FormCSSFrameworkBootstrapKey, messages.FormCSSFrameworkBootstrapValue),
			huh.NewOption(messages.FormCSSFrameworkBulmaKey, messages.FormCSSFrameworkBulmaValue),
		).
		Value(&di.Config.Frontend.CSSFramework)
}
