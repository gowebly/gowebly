package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToolVersion(t *testing.T) {
	// Test case 1: Test getting the version of a tool that is not installed.
	toolName := "tool"
	versionCommand := "--version"
	version, err := GetToolVersion(toolName, versionCommand)
	assert.Error(t, err)
	assert.EqualValues(t, "", version)

	// Test case 2: Test getting the version of a tool that is installed.
	toolName = "go"
	versionCommand = "version"
	_, err = GetToolVersion(toolName, versionCommand)
	assert.NoError(t, err)
}

func TestCheckToolIsInstalled(t *testing.T) {
	// Test case 1: Test checking if a tool is installed.
	toolName := "tool"
	versionCommand := "--version"
	isInstalled, err := CheckToolIsInstalled(toolName, versionCommand)
	assert.Error(t, err)
	assert.EqualValues(t, false, isInstalled)

	// Test case 2: Test checking if a tool is installed.
	toolName = "go"
	versionCommand = "version"
	isInstalled, err = CheckToolIsInstalled(toolName, versionCommand)
	assert.NoError(t, err)
	assert.EqualValues(t, true, isInstalled)
}
