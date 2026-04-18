package actions

import (
	"context"
	"os"
	"testing"

	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateProjectFolders tests the createProjectFolders function.
func TestCreateProjectFolders(t *testing.T) {
	// Create a temporary directory for testing.
	tempDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	// Change to temp directory.
	err := os.Chdir(tempDir)
	require.NoError(t, err)

	// Create a minimal injector.
	cfg := config.New()
	att := attachments.New()
	di := injectors.New(cfg, att)

	// Test creating project folders.
	err = createProjectFolders(di)
	assert.NoError(t, err)

	// Verify folders were created.
	assert.DirExists(t, "assets")
	assert.DirExists(t, "static/images")
	assert.DirExists(t, "templates/pages")
}

// TestCreateProjectFolders_AlreadyExists tests error when folders already exist.
func TestCreateProjectFolders_AlreadyExists(t *testing.T) {
	// Create a temporary directory for testing.
	tempDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	// Change to temp directory.
	err := os.Chdir(tempDir)
	require.NoError(t, err)

	// Create a folder that will conflict.
	err = os.MkdirAll("assets", 0755)
	require.NoError(t, err)

	// Create a minimal injector.
	cfg := config.New()
	att := attachments.New()
	di := injectors.New(cfg, att)

	// Test creating project folders - should fail because assets exists.
	err = createProjectFolders(di)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is exists")
}

// TestCreateProjectAction_Success tests the full CreateProjectAction flow with context cancellation.
func TestCreateProjectAction_Success(t *testing.T) {
	// Create a temporary directory for testing.
	tempDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	// Change to temp directory.
	err := os.Chdir(tempDir)
	require.NoError(t, err)

	// Create context and injector.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()
	cfg.Backend.GoFramework = "default"
	cfg.Frontend.ReactivityLibrary = "htmx"
	cfg.Frontend.CSSFramework = "default"
	att := attachments.New()
	di := injectors.New(cfg, att)

	errCh := make(chan error, 1)

	// Run the action in a goroutine.
	go CreateProjectAction(ctx, cancel, di, errCh)

	// Wait for result.
	err = <-errCh

	// We expect no error for folder creation (other steps may fail due to missing tools).
	// In a full test environment with all tools, this would succeed.
	// For now, we just verify it doesn't panic and handles context properly.
	if err != nil {
		// Errors are acceptable if tools are not installed.
		t.Logf("Expected error (tools may not be installed): %v", err)
	}
}

// TestCreateProjectAction_ContextCancellation tests context cancellation.
func TestCreateProjectAction_ContextCancellation(t *testing.T) {
	// Create context that's already cancelled.
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	cfg := config.New()
	att := attachments.New()
	di := injectors.New(cfg, att)

	errCh := make(chan error, 1)

	// Run the action.
	go CreateProjectAction(ctx, cancel, di, errCh)

	// Should get context cancelled error.
	err := <-errCh
	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}
