package forms

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v2/internal/fields"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/messages"
)

// RunCreateForm runs the create form.
func RunCreateForm(di *injectors.Injector) error {
	// Run welcome note form.
	if err := welcomeForm(); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "welcome", "create", err)
	}

	// Run project settings form.
	if err := projectSettingsForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "project settings", "create", err)
	}

	return nil
}

// welcomeForm runs the welcome form.
func welcomeForm() error {
	return huh.NewForm(
		huh.NewGroup(fields.WelcomeNote()), // add welcome note
	).Run()
}

// projectSettingsForm runs the project settings form.
func projectSettingsForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseAirConfirm(di)),         // confirm if Air is used
		huh.NewGroup(fields.IsUseBunConfirm(di)),         // confirm if Bun is used
		huh.NewGroup(fields.IsUseTempleConfirm(di)),      // confirm if Templ is used
		huh.NewGroup(fields.GoFrameworkSelect(di)),       // select Go framework
		huh.NewGroup(fields.ReactivityLibrarySelect(di)), // select reactivity library
		huh.NewGroup(fields.CSSFrameworkSelect(di)),      // select CSS framework
		huh.NewGroup(fields.GoModuleNameInput(di)),       // input Go module name in go.mod
		huh.NewGroup(fields.PackageNameInput(di)),        // input package name in package.json
		huh.NewGroup(fields.PortInput(di)),               // input port number
	).Run()
}
