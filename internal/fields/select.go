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
			huh.NewOptions(
				messages.FormGoFrameworkDefaultTitle,
				messages.FormGoFrameworkFiberTitle,
				messages.FormGoFrameworkGinTitle,
				messages.FormGoFrameworkEchoTitle,
				messages.FormGoFrameworkHttpRouterTitle,
			)...,
		).
		Value(&di.Config.Backend.GoFramework)
}

func ReactiveLibrarySelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormReactiveLibraryTitle).
		Description(messages.FormReactiveLibraryDescription).
		Options(
			huh.NewOptions(
				messages.FormReactiveLibraryVanillaTitle,
				messages.FormReactiveLibraryVueTitle,
				messages.FormReactiveLibraryNuxtTitle,
				messages.FormReactiveLibraryReactTitle,
				messages.FormReactiveLibraryNextTitle,
				messages.FormReactiveLibrarySvelteTitle,
				messages.FormReactiveLibrarySvelteKitTitle,
			)...,
		).
		Value(&di.Config.Frontend.ReactiveLibrary)
}

func CSSFrameworkSelect(di *injectors.Injector) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title(messages.FormCSSFrameworkTitle).
		Description(messages.FormCSSFrameworkDescription).
		Options(
			huh.NewOptions(
				messages.FormCSSFrameworkDefaultTitle,
				messages.FormCSSFrameworkTailwindCSSTitle,
				messages.FormCSSFrameworkDaisyUITitle,
				messages.FormCSSFrameworkFlowbiteTitle,
				messages.FormCSSFrameworkBootstrapTitle,
				messages.FormCSSFrameworkBulmaTitle,
			)...,
		).
		Value(&di.Config.Frontend.CSSFramework)
}
