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
		helpers.PrintStyled("Please fix error(s):", "error", "margin-top-bottom")
		helpers.PrintStyled(err.Error(), "info", "margin-left")
		helpers.PrintStyled(fmt.Sprintf(
			"For more information, see %s", constants.LinkToCompleteUserGuide),
			"warning", "margin-top-bottom",
		)
		os.Exit(1)
	}
}
