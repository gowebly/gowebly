package fields

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// WelcomeNote runs the welcome note.
func WelcomeNote() *huh.Note {
	return huh.NewNote().
		Title(fmt.Sprintf("The Gowebly CLI (%s)", variables.GoweblyVersion)).
		Description(messages.FormWelcomeDescription).
		Next(true)
}
