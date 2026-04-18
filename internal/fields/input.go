package fields

import (
	"errors"
	"regexp"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/gowebly/gowebly/v3/internal/messages"
)

// Pre-compiled regex patterns for input validation.
var (
	moduleNameRegex = regexp.MustCompile(`^[a-zA-Z0-9./_-]+$`)
	portRegex       = regexp.MustCompile(`^\d+$`)
)

// GoModuleNameInput prompts for the Go module path (e.g., github.com/user/project).
func GoModuleNameInput(di *injectors.Injector) *huh.Input {
	return huh.NewInput().
		Title(messages.FormGoModuleNameTitle).
		Description(messages.FormGoModuleNameDescription).
		Prompt(messages.FormPromptSignature).
		Validate(func(s string) error {
			if !moduleNameRegex.MatchString(s) {
				return errors.New("enter correct Go module name in the go.mod file: special characters or/and spaces are not allowed")
			}
			return nil
		}).
		Value(&di.Config.Backend.ModuleName)
}

// PortInput prompts for the backend server port number.
func PortInput(di *injectors.Injector) *huh.Input {
	return huh.NewInput().
		Title(messages.FormPortTitle).
		Description(messages.FormPortDescription).
		Prompt(messages.FormPromptSignature).
		Validate(func(s string) error {
			if !portRegex.MatchString(s) {
				return errors.New("enter correct port number: only digits allowed")
			}
			return nil
		}).
		Value(&di.Config.Backend.Port)
}

// PackageNameInput prompts for the npm package name.
func PackageNameInput(di *injectors.Injector) *huh.Input {
	return huh.NewInput().
		Title(messages.FormPackageNameTitle).
		Description(messages.FormPackageNameDescription).
		Prompt(messages.FormPromptSignature).
		Validate(func(s string) error {
			if !moduleNameRegex.MatchString(s) {
				return errors.New("enter correct project name in the package.json file: special characters or/and spaces are not allowed")
			}
			return nil
		}).
		Value(&di.Config.Frontend.PackageName)
}
