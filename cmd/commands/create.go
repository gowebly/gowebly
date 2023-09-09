package commands

import (
	"github.com/gowebly/gowebly/internal/injector"
)

// Create runs the 'gowebly create [BACKEND]' cmd command.
func Create(backend string, di *injector.Injector) error {
	// Switch between backends or create a new backend with a default.
	switch backend {
	case "fiber":
		// TODO: implement process of the creating a Go backend with Fiber to your project.
		return nil
	case "echo":
		// TODO: implement process of the creating a Go backend with Echo to your project.
		return nil
	case "chi":
		// TODO: implement process of the creating a Go backend with Chi to your project.
		return nil
	default:
		// TODO: implement process of the creating a Go backend with net/http to your project.
		return nil
	}
}
