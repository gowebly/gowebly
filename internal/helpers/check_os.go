package helpers

import "runtime"

// Check the OS and return true if it is Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
