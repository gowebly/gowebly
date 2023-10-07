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
					"Please fix error(s):", "error", "margin-top-bottom",
				},
				{
					err.Error(), "info", "margin-left",
				},
				{
					"To print all commands, just run 'gowebly' without any commands or options",
					"warning", "margin-top",
				},
				{
					fmt.Sprintf("For more information, see %s", constants.LinkToCompleteUserGuide),
					"warning", "margin-bottom",
				},
			},
		)

		// Stop with error code.
		os.Exit(1)
	}
}
