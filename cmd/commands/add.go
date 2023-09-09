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
	case "unocss":
		//
		return nil
	case "tailwindcss":
		//
		return nil
	default:
		//
		return errors.New(constants.ErrorRunAddCommandWithUnknownCSSFrameworkName)
	}
}
