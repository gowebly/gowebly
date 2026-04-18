package helpers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gowebly/gowebly/v3/internal/messages"
)

// GetToolVersion executes the tool with the given version command and returns its output.
func GetToolVersion(name, versionCommand string) (string, error) {
	cmd := exec.Command(name, versionCommand)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(messages.ErrorCMDNotExecuteCommand, name, versionCommand, err)
	}

	return strings.TrimSpace(string(output)), nil
}

// CheckToolIsInstalled verifies if a tool exists in PATH by running the version command.
func CheckToolIsInstalled(name, versionCommand string) (bool, error) {
	cmd := exec.Command(name, versionCommand)

	_, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(messages.ErrorCMDNotExecuteCommand, name, versionCommand, err)
	}

	return true, nil
}

// GetToolVersionSafe gets the version of a tool without returning an error.
// Returns the version string and a boolean indicating if the tool is installed.
// This is useful for diagnostic commands where missing tools are expected.
func GetToolVersionSafe(name, versionCommand string) (string, bool) {
	version, err := GetToolVersion(name, versionCommand)
	if err != nil {
		return "", false
	}
	return version, true
}
