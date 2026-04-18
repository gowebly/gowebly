package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gowebly/gowebly/v3/internal/messages"
)

// File represents a file with name and content data.
type File struct {
	Name string
	Data []byte
}

// IsExistInFolder checks if a file or folder exists at the given path.
func IsExistInFolder(name string, isFolder bool) bool {
	info, err := os.Stat(filepath.Clean(name))
	if err == nil && info != nil {
		return info.IsDir() == isFolder
	}

	return false
}

// MakeFile creates a single file with the given name and data.
func MakeFile(file File) error {
	if IsExistInFolder(file.Name, false) {
		return fmt.Errorf(messages.ErrorOSFileIsExists, file.Name)
	}

	return os.WriteFile(file.Name, file.Data, 0o600)
}

// MakeFiles creates multiple files from the given slice.
func MakeFiles(files []File) error {
	for _, f := range files {
		if err := MakeFile(f); err != nil {
			return err
		}
	}

	return nil
}

// MakeFolders creates multiple directories with the given names.
func MakeFolders(names ...string) error {
	for _, name := range names {
		if IsExistInFolder(name, true) {
			return fmt.Errorf(messages.ErrorOSFolderIsExists, name)
		}

		if err := os.MkdirAll(name, 0o700); err != nil {
			return err
		}
	}

	return nil
}
