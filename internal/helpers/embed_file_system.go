package helpers

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gowebly/gowebly/v3/internal/messages"
)

// EmbedFile represents a file to be copied from embedded filesystem.
type EmbedFile struct {
	EmbedFile  string
	OutputFile string
}

// EmbedTemplate represents a template file with data to be generated.
type EmbedTemplate struct {
	EmbedFile  string
	OutputFile string
	Data       any
}

// CopyFilesFromEmbedFS copies files from embedded filesystem to output directory.
func CopyFilesFromEmbedFS(efs embed.FS, files []EmbedFile) error {
	for _, f := range files {
		data, err := efs.ReadFile(f.EmbedFile)
		if err != nil {
			return err
		}

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

// GenerateFilesByTemplateFromEmbedFS processes templates and writes generated files to output.
// Uses a temporary directory during processing to avoid partial writes.
func GenerateFilesByTemplateFromEmbedFS(efs embed.FS, templates []EmbedTemplate) error {
	tempFolder, err := os.MkdirTemp("", "*")
	if err != nil {
		return err
	}

	for _, t := range templates {
		if IsExistInFolder(t.OutputFile, false) {
			return fmt.Errorf(messages.ErrorOSFileIsExists, t.OutputFile)
		}

		tmpl, err := template.New(filepath.Base(t.EmbedFile)).
			Funcs(template.FuncMap{
				"trim": strings.TrimSpace,
			}).
			ParseFS(efs, t.EmbedFile)
		if err != nil {
			return err
		}

		tempFile, err := os.CreateTemp(tempFolder, "*")
		if err != nil {
			return err
		}

		if err := tmpl.Execute(tempFile, t.Data); err != nil {
			return err
		}

		if _, err := tempFile.Seek(0, 0); err != nil {
			return err
		}

		outputFile, err := os.Create(t.OutputFile)
		if err != nil {
			return err
		}

		if _, err := io.Copy(outputFile, tempFile); err != nil {
			return err
		}

		if err := tempFile.Close(); err != nil {
			return err
		}

		if err := outputFile.Close(); err != nil {
			return err
		}
	}

	return os.RemoveAll(tempFolder)
}
