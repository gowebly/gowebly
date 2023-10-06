package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gowebly/gowebly/internal/constants"
)

// File represent a struct to a file.
type File struct {
	Name string
	Data []byte
}

// IsExistInFolder searches for a file or folder by the given name in the
// current folder.
func IsExistInFolder(name string, isFolder bool) bool {
	// Check, if file or folder is existing.
	info, err := os.Stat(filepath.Clean(name))
	if err == nil || !os.IsNotExist(err) {
		return info.IsDir() == isFolder
	}

	return false
}

// MakeFile makes a single file with name and data.
func MakeFile(file File) error {
	// Check, if file is existing.
	if IsExistInFolder(file.Name, false) {
		return fmt.Errorf(constants.ErrorOSFileIsExists, file.Name)
	}

	return os.WriteFile(file.Name, file.Data, 0o600)
}

// MakeFiles makes a multiply files with names and data.
func MakeFiles(files []File) error {
	for _, f := range files {
		// Create a new file.
		if err := MakeFile(f); err != nil {
			return err
		}
	}

	return nil
}

// MakeFolders makes a multiply folders with names.
func MakeFolders(names ...string) error {
	for _, name := range names {
		// Check, if folder is existing.
		if IsExistInFolder(name, true) {
			return fmt.Errorf(constants.ErrorOSFolderIsExists, name)
		}

		// Create a new folder.
		if err := os.MkdirAll(name, 0o700); err != nil {
			return err
		}
	}

	return nil
}

// RemoveFiles removes files by names.
func RemoveFiles(names ...string) error {
	for _, name := range names {
		// Remove file by name.
		if err := os.Remove(name); err != nil {
			return fmt.Errorf(constants.ErrorOSRemoveFile, name)
		}
	}

	return nil
}
