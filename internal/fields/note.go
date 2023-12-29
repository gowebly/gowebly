package fields

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/internal/messages"
	"github.com/gowebly/gowebly/internal/variables"
)

func WelcomeNote() *huh.Note {
	return huh.NewNote().
		Title(fmt.Sprintf("The Gowebly CLI (%s)", variables.GoweblyVersion)).
		Description(messages.FormWelcomeDescription).
		Next(true)
}
