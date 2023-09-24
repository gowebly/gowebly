package helpers

import (
	"embed"
	"errors"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
)

// EmbedFile represent struct for embed file system.
type EmbedFile struct {
	EmbedFolder  string
	OutputFolder string
}

// EmbedTemplate represents struct for a one template.
type EmbedTemplate struct {
	EmbedFile    string
	OutputFile   string
	OutputFolder string
	Data         any
}

// CopyFromEmbedFS function for copy files from the embed file system.
func CopyFromEmbedFS(efs embed.FS, files []EmbedFile) error {
	for _, f := range files {
		// Return copied folders and files.
		if err := fs.WalkDir(efs, f.EmbedFolder, func(path string, entry fs.DirEntry, err error) error {
			// Checking embed path.
			if err != nil {
				return errors.New("os: can't copy files from embed file system")
			}

			// Check, if file is existing.
			if IsExistInFolder(filepath.Join(f.OutputFolder, entry.Name()), false) {
				return errors.New(constants.ErrorProjectFolderIsNotEmpty)
			}

			// Checking, if embedded file is not a folder.
			if !entry.IsDir() {
				// Set file data.
				fileData, errReadFile := fs.ReadFile(efs, path)
				if errReadFile != nil {
					return errReadFile
				}

				// Create file from embedded.
				if errMakeFile := MakeFile(filepath.Join(f.OutputFolder, entry.Name()), fileData); errMakeFile != nil {
					return errMakeFile
				}
			}

			return nil
		}); err != nil {
			return err
		}
	}

	return nil
}

// GenerateFromEmbedFS generates new files from the given template with
// variables from embed FS.
func GenerateFromEmbedFS(efs embed.FS, templates []EmbedTemplate) error {
	for _, t := range templates {
		// Check, if file is existing.
		if IsExistInFolder(filepath.Join(t.OutputFolder, t.OutputFile), false) {
			return errors.New(constants.ErrorProjectFolderIsNotEmpty)
		}

		// Parse template from embed file system.
		tmpl, err := template.ParseFS(efs, t.EmbedFile)
		if err != nil {
			return err
		}

		// Create a new temp file with the given data.
		file, err := os.CreateTemp("", "*")
		if err != nil {
			return err
		}

		// Rename temp file.
		if err := os.Rename(file.Name(), filepath.Join(t.OutputFolder, t.OutputFile)); err != nil {
			return err
		}

		// Set variables to the given.
		if err := tmpl.Execute(file, t.Data); err != nil {
			return err
		}

		// Close file.
		if err := file.Close(); err != nil {
			return err
		}
	}

	return nil
}
