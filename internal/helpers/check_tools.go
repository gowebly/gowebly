package helpers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gowebly/gowebly/internal/messages"
)

// GetToolVersion gets the version of a tool by executing a command.
//
// It takes the name of the tool and the command to retrieve the version as input.
// It returns the version as a string and any error encountered.
func GetToolVersion(name, versionCommand string) (string, error) {
	// Create a new command with the given tool name and version command.
	cmd := exec.Command(name, versionCommand)

	// Execute the command and retrieve the output.
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(messages.ErrorCMDNotExecuteCommand, name, versionCommand, err)
	}

	// Return the output as a string and nil error.
	return strings.TrimSpace(string(output)), nil
}

// CheckToolIsInstalled checks if a tool is installed by executing a command.
//
// It takes the name of the tool and the command to retrieve the version as input.
// It returns the bool indicating if the tool is installed and any error encountered.
func CheckToolIsInstalled(name, versionCommand string) (bool, error) {
	// Create a new command with the given tool name and version command.
	cmd := exec.Command(name, versionCommand)

	// Execute the command and retrieve the output.
	_, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(messages.ErrorCMDNotExecuteCommand, name, versionCommand, err)
	}

	// Return the output as a string and nil error.
	return true, nil
}
