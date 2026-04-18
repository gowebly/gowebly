package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
)

// IsUseAirConfirm prompts for Air live-reload tool installation.
func IsUseAirConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormAirUsageTitle).
		Description(messages.FormAirUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseAir)
}

// IsUseTempleConfirm prompts for Templ HTML template engine installation.
func IsUseTempleConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormTemplUsageTitle).
		Description(messages.FormTemplUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseTempl)
}

// IsUseBunConfirm prompts for Bun JavaScript runtime selection.
func IsUseBunConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormBunUsageTitle).
		Description(messages.FormBunUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseBun)
}

// IsUseGolangCILintConfirm prompts for golangci-lint installation.
func IsUseGolangCILintConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormGolangCILintUsageTitle).
		Description(messages.FormGolangCILintUsageDescription).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseGolangCILint)
}
