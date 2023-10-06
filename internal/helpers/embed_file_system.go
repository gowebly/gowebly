package helpers

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/gowebly/gowebly/internal/constants"
)

// EmbedFile represent struct for embed file system.
type EmbedFile struct {
	EmbedFile  string
	OutputFile string
}

// EmbedTemplate represents struct for a one template.
type EmbedTemplate struct {
	EmbedFile  string
	OutputFile string
	Data       any
}

// CopyFilesFromEmbedFS function for copy all files from the embed file system.
func CopyFilesFromEmbedFS(efs embed.FS, files []EmbedFile) error {
	for _, f := range files {
		// Read the attachments (embedded) config file.
		data, err := efs.ReadFile(f.EmbedFile)
		if err != nil {
			return err
		}

		// Create a config file.
		if err := MakeFiles(
			[]File{
				{
					f.OutputFile, data,
				},
			},
		); err != nil {
			return err
		}
	}

	return nil
}

// GenerateFilesByTemplateFromEmbedFS generates new files from the given template with
// variables from embed FS.
func GenerateFilesByTemplateFromEmbedFS(efs embed.FS, templates []EmbedTemplate) error {
	// Create a new temp folder.
	tempFolder, err := os.MkdirTemp("", "*")
	if err != nil {
		return err
	}

	for _, t := range templates {
		// Check if file exists.
		if IsExistInFolder(t.OutputFile, false) {
			return fmt.Errorf(constants.ErrorOSFileIsExists, t.OutputFile)
		}

		// Parse template from embed file system.
		tmpl, err := template.ParseFS(efs, t.EmbedFile)
		if err != nil {
			return err
		}

		// Create a new temp file with the given data.
		tempFile, err := os.CreateTemp(tempFolder, "*")
		if err != nil {
			return err
		}

		// Set variables to the given.
		if err := tmpl.Execute(tempFile, t.Data); err != nil {
			return err
		}

		// Reset the record position to the beginning of the file.
		if _, err := tempFile.Seek(0, 0); err != nil {
			return err
		}

		// Copy temp file to the output file.
		outputFile, err := os.Create(t.OutputFile)
		if err != nil {
			return err
		}

		// Copy file from the temp file to the output.
		if _, err := io.Copy(outputFile, tempFile); err != nil {
			return err
		}

		// Close temp file.
		if err := tempFile.Close(); err != nil {
			return err
		}

		// Close output file.
		if err := outputFile.Close(); err != nil {
			return err
		}
	}

	return os.RemoveAll(tempFolder)
}
