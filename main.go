package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gowebly/gowebly/v2/cmd"
	"github.com/gowebly/gowebly/v2/internal/helpers"
	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/gowebly/gowebly/v2/internal/variables"
)

func main() {
	// Parse flags.
	flag.Parse()

	// Run cmd with the parsed flags or exit with error.
	if err := cmd.Run(flag.Args()); err != nil {
		// Show error.
		fmt.Println(helpers.MakeStyled(
			messages.CommandErrorSummaryTitle,
			&helpers.StringStyle{Color: variables.ColorRed, IsBold: true},
		))
		fmt.Println(helpers.MakeStyledFrame(
			err.Error(),
			&helpers.FrameStyle{Padding: []int{1}, Color: variables.ColorRed},
		))
		fmt.Println(helpers.MakeStyled(
			messages.CommandMoreInformationTitle,
			&helpers.StringStyle{Color: variables.ColorYellow},
		))

		// Stop with error code.
		os.Exit(1)
	}
}
