package forms

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/internal/fields"
	"github.com/gowebly/gowebly/internal/injectors"
	"github.com/gowebly/gowebly/internal/messages"
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

	// Run Go framework form.
	if err := goFrameworkForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "go framework", "create", err)
	}

	// Run HTMX form.
	if err := htmxForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "htmx", "create", err)
	}

	// Check, if HTMX is used.
	if di.Config.Frontend.IsUseHTMX {
		// If yes, run Templ form.
		if err := templForm(di); err != nil {
			return fmt.Errorf(messages.ErrorFormNotRun, "templ", "create", err)
		}
	} else {
		// If not, run reactive library form.
		if err := reactiveLibraryForm(di); err != nil {
			return fmt.Errorf(messages.ErrorFormNotRun, "reactive library", "create", err)
		}
	}

	// Run CSS framework form.
	if err := cssFrameworkForm(di); err != nil {
		return fmt.Errorf(messages.ErrorFormNotRun, "css framework", "create", err)
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
		huh.NewGroup(fields.IsUseAirConfirm(di)),   // confirm if Air is used
		huh.NewGroup(fields.GoModuleNameInput(di)), // input Go module name in go.mod
		huh.NewGroup(fields.PortInput(di)),         // input port number
		huh.NewGroup(fields.IsUseBunConfirm(di)),   // confirm if Bun is used
		huh.NewGroup(fields.PackageNameInput(di)),  // input package name in package.json
	).Run()
}

// goFrameworkForm runs the Go framework form.
func goFrameworkForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.GoFrameworkSelect(di)), // select Go framework
	).Run()
}

// htmxForm runs the HTMX form.
func htmxForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseHTMXConfirm(di)), // confirm if htmx is used
	).Run()
}

// templForm runs the Templ form.
func templForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseTempleConfirm(di)), // confirm if Templ is used
	).Run()
}

// reactiveLibraryForm runs the reactive library form.
func reactiveLibraryForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.ReactiveLibrarySelect(di)), // select reactive library
	).Run()
}

// cssFrameworkForm runs the CSS framework form.
func cssFrameworkForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.CSSFrameworkSelect(di)), // select CSS framework
	).Run()
}
