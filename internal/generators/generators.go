package generators

import (
	"embed"
	"errors"
	"html/template"
	"os"

	"github.com/gowebly/gowebly/internal/constants"
	"github.com/gowebly/gowebly/internal/helpers"
)

type Template struct {
	Name   string
	Output string
	Data   any
}

// FilesFromEmbedFS generates new files from the given template with
// variables from embed FS.
func FilesFromEmbedFS(fs embed.FS, templates []Template) error {
	for _, t := range templates {
		// Check, if file is existing.
		if helpers.IsExistInCurrentFolder(t.Output, false) {
			return errors.New(constants.ErrorProjectFolderIsNotEmpty)
		}

		// Parse template from embed file system.
		tmpl, err := template.ParseFS(fs, t.Name)
		if err != nil {
			return err
		}

		// Create a new temp file with the given data.
		file, err := os.CreateTemp("", "*")
		if err != nil {
			return err
		}

		// Rename temp file.
		if err = os.Rename(file.Name(), t.Output); err != nil {
			return err
		}

		// Set variables to the given.
		if err = tmpl.Execute(file, t.Data); err != nil {
			return err
		}

		// Close file.
		if err = file.Close(); err != nil {
			return err
		}
	}

	return nil
}
