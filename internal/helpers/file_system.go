package helpers

import (
	"fmt"
	"os"
	"path/filepath"
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

// MakeFiles makes a multiply files with names and data.
func MakeFiles(files []File) error {
	for _, f := range files {
		// Check, if file is existing.
		if IsExistInFolder(f.Name, false) {
			return fmt.Errorf("os: file with name '%s' is found in the current folder, cannot be overwritten", f.Name)
		}

		// Create a new file.
		if err := os.WriteFile(f.Name, f.Data, 0o600); err != nil {
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
