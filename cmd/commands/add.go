package commands

import (
	"errors"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/injector"
)

// Add runs the 'gowebly add [CSS_FRAMEWORK]' cmd command.
func Add(cssFramework string, di *injector.Injector) error {
	// Switch between CSS frameworks or return error.
	switch cssFramework {
	case "tailwindcss":
		// TODO: implement process of the adding Tailwind CSS framework to your project.
		return nil
	case "unocss":
		// TODO: implement process of the adding UnoCSS framework to your project.
		return nil
	default:
		// Returning error message.
		return errors.New(constants.ErrorRunAddCommandWithUnknownCSSFrameworkName)
	}
}
