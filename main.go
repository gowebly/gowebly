package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/gowebly/gowebly/cmd"
	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
)

func main() {
	// Parse flags.
	flag.Parse()

	// Inject all dependencies (config, embed files).
	di, err := inject(filepath.Clean(constants.YAMLConfigFileName))
	if err != nil {
		helpers.PrintStyled(err.Error(), "error", "")
		os.Exit(1)
	}

	// Run cmd with the parsed flags or exit with error.
	if err = cmd.Run(flag.Args(), di); err != nil {
		helpers.PrintStyled(err.Error(), "error", "")
		os.Exit(1)
	}
}
