package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/internal/fields"
	"github.com/gowebly/gowebly/internal/injectors"
)

// CreateWelcomeForm runs the welcome form.
func CreateWelcomeForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.WelcomeNote()), // add welcome note
	).Run()
}

// CreateProjectSettingsForm runs the project settings form.
func CreateProjectSettingsForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.GoModuleNameInput(di)), // input Go module name in go.mod
		huh.NewGroup(fields.IsUseAirConfirm(di)),   // confirm if Air is used
		huh.NewGroup(fields.PackageNameInput(di)),  // input package name in package.json
		huh.NewGroup(fields.IsUseBunConfirm(di)),   // confirm if Bun is used
	).Run()
}

// CreateGoFrameworkForm runs the Go framework form.
func CreateGoFrameworkForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.GoFrameworkSelect(di)), // select Go framework
	).Run()
}

// CreateHTMXForm runs the HTMX form.
func CreateHTMXForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseHTMXConfirm(di)), // confirm if htmx is used
	).Run()
}

// CreateTemplForm runs the Templ form.
func CreateTemplForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseTempleConfirm(di)), // confirm if Templ is used
	).Run()
}

// CreateHTMXForm runs the HTMX form.
func CreateHyperscriptForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.IsUseHyperscriptConfirm(di)), // confirm if hyperscript is used
	).Run()
}

// CreateReactiveLibraryForm runs the reactive library form.
func CreateReactiveLibraryForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.ReactiveLibrarySelect(di)), // select reactive library
	).Run()
}

// CreateCSSFrameworkForm runs the CSS framework form.
func CreateCSSFrameworkForm(di *injectors.Injector) error {
	return huh.NewForm(
		huh.NewGroup(fields.CSSFrameworkSelect(di)), // select CSS framework
	).Run()
}
