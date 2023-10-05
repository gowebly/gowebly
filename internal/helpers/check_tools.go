package helpers

import (
	"fmt"
	"os/exec"
	"strings"
)

// Tool represents a struct for a CLI tool.
type Tool struct {
	Name, VersionCommand string
}

// CheckTools checks a list with tools and collect its versions.
func CheckTools(tools []Tool) map[string][]string {
	// Create a new map.
	results := make(map[string][]string, 0)

	for _, t := range tools {
		// Create a new cmd execution with output.
		output, err := exec.Command(t.Name, t.VersionCommand).Output()
		if err != nil {
			// If the current tool is not exists, save error message to the map.
			results[t.Name] = []string{
				fmt.Sprintf("'%s' tool is not installed", t.Name),
				"error",
			}
		} else {
			// Else, save the current tool version to the map.
			results[t.Name] = []string{
				strings.ToLower(strings.Trim(string(output), "\n")),
				"success",
			}
		}
	}

	return results
}
