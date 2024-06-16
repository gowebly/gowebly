package fields

import (
	"errors"
	"regexp"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v2/internal/injectors"
	"github.com/gowebly/gowebly/v2/internal/messages"
)

// GoModuleNameInput runs the go module name input.
func GoModuleNameInput(di *injectors.Injector) *huh.Input {
	return huh.NewInput().
		Title(messages.FormGoModuleNameTitle).
		Description(messages.FormGoModuleNameDescription).
		Prompt(messages.FormPromptSignature).
		Validate(func(s string) error {
			if !regexp.MustCompile(`^[a-zA-Z0-9./_-]+$`).MatchString(s) {
				return errors.New("enter correct Go module name in the go.mod file: special characters or/and spaces are not allowed")
			}
			return nil
		}).
		Value(&di.Config.Backend.ModuleName)
}

// PortInput runs the port input.
func PortInput(di *injectors.Injector) *huh.Input {
	return huh.NewInput().
		Title(messages.FormPortTitle).
		Description(messages.FormPortDescription).
		Prompt(messages.FormPromptSignature).
		Validate(func(s string) error {
			if !regexp.MustCompile(`^\d+$`).MatchString(s) {
				return errors.New("enter correct port number: only digits allowed")
			}
			return nil
		}).
		Value(&di.Config.Backend.Port)
}

// PackageNameInput runs the package name input.
func PackageNameInput(di *injectors.Injector) *huh.Input {
	return huh.NewInput().
		Title(messages.FormPackageNameTitle).
		Description(messages.FormPackageNameDescription).
		Prompt(messages.FormPromptSignature).
		Validate(func(s string) error {
			if !regexp.MustCompile(`^[a-zA-Z0-9./_-]+$`).MatchString(s) {
				return errors.New("enter correct project name in the package.json file: special characters or/and spaces are not allowed")
			}
			return nil
		}).
		Value(&di.Config.Frontend.PackageName)
}
