package fields

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/gowebly/gowebly/v3/internal/messages"
	"github.com/gowebly/gowebly/v3/internal/variables"
)

// WelcomeNote displays the CLI welcome message with version info.
func WelcomeNote() *huh.Note {
	return huh.NewNote().
		Title(fmt.Sprintf("The Gowebly CLI (%s)", variables.GoweblyVersion)).
		Description(messages.FormWelcomeDescription)
}
