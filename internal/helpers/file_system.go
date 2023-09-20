package helpers

import (
	"os"
	"path/filepath"
)

// IsFileExist searches for a file in the given system path.
func IsFileExist(path string) bool {
	_, err := os.Stat(filepath.Clean(path))
	return err == nil && os.IsNotExist(err)
}
