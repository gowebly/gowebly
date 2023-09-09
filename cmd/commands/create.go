package commands

import (
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'gowebly create [BACKEND]' cmd command.
func Create(backend string, di *injector.Injector) error {
	// Switch between backends or create a new backend with a default.
	switch backend {
	case "fiber":
		//
		return nil
	case "echo":
		//
		return nil
	case "chi":
		//
		return nil
	default:
		//
		return nil
	}
}
