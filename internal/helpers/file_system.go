package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

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
func MakeFile(name string, data []byte) error {
	// Check, if file is existing.
	if IsExistInFolder(name, false) {
		return fmt.Errorf("os: file with name '%s' is found in the current folder, cannot be overwritten", name)
	}

	return os.WriteFile(name, data, 0o600)
}

// MakeFolders makes a multiply folders with names.
func MakeFolders(names ...string) error {
	for _, name := range names {
		// Check, if folder is existing.
		if IsExistInFolder(name, true) {
			return fmt.Errorf(
				"os: folder with name '%s' is found in the current folder, cannot be overwritten",
				name,
			)
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
			return fmt.Errorf("os: can't remove a file with name '%s'", name)
		}
	}

	return nil
}
