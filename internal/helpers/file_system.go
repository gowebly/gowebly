package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsExistInCurrentFolder searches for a file or folder by the given name in the
// current folder.
func IsExistInCurrentFolder(name string, isFolder bool) bool {
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
	if IsExistInCurrentFolder(name, false) {
		return fmt.Errorf("a file with name '%s' is found in the current folder, cannot be overwritten", name)
	}

	return os.WriteFile(name, data, 0o700)
}

// MakeFolder makes a single folder with name.
func MakeFolder(name string) error {
	// Check, if folder is existing.
	if IsExistInCurrentFolder(name, true) {
		return fmt.Errorf("a folder with name '%s' is found in the current folder, cannot be overwritten", name)
	}

	return os.Mkdir(name, 0o700)
}

// MakeFolders makes a multiply folders with names.
func MakeFolders(names ...string) error {
	for _, name := range names {
		// Check, if folder is existing.
		if IsExistInCurrentFolder(name, true) {
			return fmt.Errorf("a folder with name '%s' is found in the current folder, cannot be overwritten", name)
		}

		// Create a new folder.
		if err := os.MkdirAll(name, 0o700); err != nil {
			return err
		}
	}

	return nil
}

// RemoveFolders removes folders by names with all files.
func RemoveFolders(names []string) error {
	// Remove folders in the loop.
	for _, name := range names {
		if err := os.RemoveAll(name); err != nil {
			return fmt.Errorf("can't remove a folder with name '%s' from the current folder", name)
		}
	}

	return nil
}
