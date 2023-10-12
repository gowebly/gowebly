package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gowebly/gowebly/cmd"
	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
)

func main() {
	// Parse flags.
	flag.Parse()

	// Run cmd with the parsed flags or exit with error.
	if err := cmd.Run(flag.Args()); err != nil {
		// Print block of messages.
		helpers.PrintStyledBlock(
			[]helpers.StyledOutput{
				{
					Text: "Please fix error(s):", State: "error", Style: "margin-top-bottom",
				},
				{
					Text: err.Error(), State: "info", Style: "margin-left",
				},
				{
					Text:  "To print all commands, just run 'gowebly' without any commands or options",
					State: "warning", Style: "margin-top",
				},
				{
					Text:  fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
					State: "warning", Style: "margin-bottom",
				},
			},
		)

		// Stop with error code.
		os.Exit(1)
	}
}
