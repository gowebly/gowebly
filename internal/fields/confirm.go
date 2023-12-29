package fields

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
)

func IsUseAirConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormAirUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseAir)
}

func IsUseTempleConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormTemplUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseTempl)
}

func IsUseHTMXConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormHTMXUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Frontend.IsUseHTMX)
}

func IsUseBunConfirm(di *injectors.Injector) *huh.Confirm {
	return huh.NewConfirm().
		Title(messages.FormBunUsageTitle).
		Affirmative("Yes").
		Negative("No").
		Value(&di.Config.Tools.IsUseBun)
}
