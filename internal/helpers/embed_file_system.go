package helpers

import (
	"embed"
	"errors"
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
		if err := renameFile(file.Name(), t.OutputFile); err != nil {
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

func renameFile(src string, dst string) (err error) {
	err = copyFile(src, dst)
	if err != nil {
		return fmt.Errorf("failed to copy source file %s to %s: %s", src, dst, err)
	}
	err = os.RemoveAll(src)
	if err != nil {
		return fmt.Errorf("failed to cleanup source file %s: %s", src, err)
	}
	return nil
}

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}
