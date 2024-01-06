package commands

import (
	"fmt"

	"github.com/gowebly/gowebly/v2/internal/helpers"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

// Unknown runs an unknown cmd command or nothing.
func Unknown() error {
	// Add technical space for 'unknown' command.
	fmt.Println()

	// Generate content body.
	contentBody := fmt.Sprintf(
		messages.CommandUnknownSummaryDescription,
		helpers.MakeStyled(messages.CommandUnknownSummarySubTitle, &helpers.StringStyle{Color: variables.ColorGrey}),
		helpers.MakeStyled("create", &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled("doctor", &helpers.StringStyle{Color: variables.ColorBlue}),
		helpers.MakeStyled("run", &helpers.StringStyle{Color: variables.ColorBlue}),
	)

	// Show created project info.
	fmt.Println(helpers.MakeStyled(
		messages.CommandUnknownSummaryTitle,
		&helpers.StringStyle{Color: variables.ColorBlue, IsBold: true},
	))
	fmt.Println(helpers.MakeStyledFrame(
		contentBody,
		&helpers.FrameStyle{Padding: []int{1}, Color: variables.ColorBlue},
	))
	fmt.Println(helpers.MakeStyled(
		messages.CommandMoreInformationTitle,
		&helpers.StringStyle{Color: variables.ColorYellow},
	))

	return nil
}
