package forms

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v3/internal/fields"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
)

// RunCreateForm displays the interactive project creation form wizard.
func RunCreateForm(di *injectors.Injector) error {
	if err := welcomeForm(); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "welcome", "create", err)
	}

	if err := projectSettingsForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "project settings", "create", err)
	}

	return nil
}

func welcomeForm() error {
	return huh.NewForm(
		huh.NewGroup(fields.WelcomeNote()),
	).Run()
}

func projectSettingsForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseAirConfirm(di)),
		huh.NewGroup(fields.IsUseBunConfirm(di)),
		huh.NewGroup(fields.IsUseTempleConfirm(di)),
		huh.NewGroup(fields.IsUseGolangCILintConfirm(di)),
		huh.NewGroup(fields.GoFrameworkSelect(di)),
		huh.NewGroup(fields.ReactivityLibrarySelect(di)),
		huh.NewGroup(fields.CSSFrameworkSelect(di)),
		huh.NewGroup(fields.GoModuleNameInput(di)),
		huh.NewGroup(fields.PackageNameInput(di)),
		huh.NewGroup(fields.PortInput(di)),
	).Run()
}
