package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
)

// IsUseAirConfirm runs the confirm if Air is used.
func IsUseAirConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormAirUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseAir)
}

// IsUseTempleConfirm runs the confirm if Templ is used.
func IsUseTempleConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormTemplUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseTempl)
}

// IsUseBunConfirm runs the confirm if Bun is used.
func IsUseBunConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormBunUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseBun)
}
