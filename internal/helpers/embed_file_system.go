package helpers

import (
	"embed"
	"errors"
	"html/template"
	"io"
	"os"
	"path/filepath"

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
		if err := MakeFiles([]File{
			{
				f.OutputFile, data,
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

// GenerateFilesByTemplateFromEmbedFS generates new files from the given template with
// variables from embed FS.
func GenerateFilesByTemplateFromEmbedFS(efs embed.FS, templates []EmbedTemplate) error {
	for _, t := range templates {
		// Check, if file is existing.
		if IsExistInFolder(t.OutputFile, false) {
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
		if err := os.Rename(file.Name(), t.OutputFile); err != nil {
			// Define a new variable for the exit error.
			var exitErr *os.LinkError

			// Check, if the current error is exit error.
			if errors.As(err, &exitErr) {
				// Try to rename file in non-Unix system.
				if err := renameFileNonUnix(file.Name(), t.OutputFile); err != nil {
					return err
				}
			} else {
				return err
			}
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

// renameFileNonUnix renames the given src file name to the destination for the
// non-Unix systems.
func renameFileNonUnix(src, dest string) error {
	// Open the source file.
	input, err := os.Open(filepath.Clean(src))
	if err != nil {
		return err
	}
	defer input.Close()

	// Create the destination file.
	output, err := os.Create(filepath.Clean(dest))
	if err != nil {
		return err
	}
	defer output.Close()

	// Copy all content of the source file to the destination.
	_, err = io.Copy(output, input)
	if err != nil {
		return err
	}

	// Commit the current contents of the destination file to stable storage.
	if err := output.Sync(); err != nil {
		return err
	}

	// Get file info from the source file.
	inputFileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Set the right chmod to the destination file.
	if err := os.Chmod(dest, inputFileInfo.Mode()); err != nil {
		return err
	}

	// Cleanup the source file.
	if err := os.RemoveAll(src); err != nil {
		return err
	}

	return nil
}
