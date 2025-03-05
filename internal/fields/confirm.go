package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
)

// IsUseAirConfirm runs the confirm if Air is used.
func IsUseAirConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormAirUsageTitle).
		Description(messages.FormAirUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseAir)
}

// IsUseTempleConfirm runs the confirm if Templ is used.
func IsUseTempleConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormTemplUsageTitle).
		Description(messages.FormTemplUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseTempl)
}

// IsUseBunConfirm runs the confirm if Bun is used.
func IsUseBunConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormBunUsageTitle).
		Description(messages.FormBunUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseBun)
}

// IsUseGolangCILintConfirm runs the confirm if Golang CI Lint is used.
func IsUseGolangCILintConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormGolangCILintUsageTitle).
		Description(messages.FormGolangCILintUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseGolangCILint)
}
