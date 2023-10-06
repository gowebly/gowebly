package helpers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gowebly/gowebly/internal/constants"
)

// Tool represents a struct for a tool.
type Tool struct {
	Name, VersionCommand string
}

// resultsOutput represents a (private) struct for results output.
type resultsOutput struct {
	Status, Output string
}

// CheckTools checks a list with tools and collect its versions.
func CheckTools(tools []Tool) []resultsOutput {
	// Create a new slice for results.
	results := make([]resultsOutput, 0)

	for _, t := range tools {
		// Create a new cmd execution with output.
		output, err := exec.Command(t.Name, t.VersionCommand).Output()
		if err != nil {
			// If the current tool is not exists, save error message to the map.
			results = append(results, resultsOutput{
				Status: "error",
				Output: fmt.Sprintf(constants.ErrorHelperToolNotInstalled, t.Name),
			})
		} else {
			// Else, save the current tool version to the map.
			results = append(results, resultsOutput{
				Status: "success",
				Output: strings.ToLower(strings.Trim(string(output), "\n")),
			})
		}
	}

	return results
}
