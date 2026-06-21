package actions

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/gowebly/gowebly/v3/internal/attachments"
	"github.com/gowebly/gowebly/v3/internal/config"
	"github.com/gowebly/gowebly/v3/internal/helpers"
	"github.com/gowebly/gowebly/v3/internal/injectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generatePocketsbaseBackend generates only the backend files for the PocketBase
// framework into the given directory using the provided config.
func generatePocketbaseBackend(t *testing.T, dir string, useTempl bool) {
	t.Helper()

	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	require.NoError(t, os.Chdir(dir))


	cfg := config.New()
	cfg.Backend.GoFramework = "pocketbase"
	cfg.Backend.ModuleName = "github.com/test/mfp"
	cfg.Backend.Port = "8080"
	cfg.Tools.IsUseTempl = useTempl
	att := attachments.New()
	di := injectors.New(cfg, att)

	require.NoError(t, createProjectFolders(di))
	require.NoError(t, createBackendFiles(di))

	if useTempl {
		// Generate templ files from frontend templates.
		templates := []helpers.EmbedTemplate{
			{
				EmbedFile:  "templates/frontend/main.templ.gotmpl",
				OutputFile: "templates/main.templ",
				Data:       cfg,
			},
			{
				EmbedFile:  "templates/frontend/index.templ.gotmpl",
				OutputFile: "templates/pages/index.templ",
				Data:       cfg,
			},
		}
		require.NoError(t, helpers.GenerateFilesByTemplateFromEmbedFS(att.Templates, templates))
	} else {
		// Generate html files from frontend templates.
		templates := []helpers.EmbedTemplate{
			{
				EmbedFile:  "templates/frontend/main.html.gotmpl",
				OutputFile: "templates/main.html",
				Data:       cfg,
			},
			{
				EmbedFile:  "templates/frontend/index.html.gotmpl",
				OutputFile: "templates/pages/index.html",
				Data:       cfg,
			},
		}
		require.NoError(t, helpers.GenerateFilesByTemplateFromEmbedFS(att.Templates, templates))
	}
}

// TestPocketbaseTemplatesGenerateAndCompile_Templ verifies that the generated
// PocketBase backend code (templ variant) compiles successfully.
func TestPocketbaseTemplatesGenerateAndCompile_Templ(t *testing.T) {
	// integration test - requires templ CLI

	dir := t.TempDir()
	generatePocketbaseBackend(t, dir, true)

	// Run templ generate.
	if err := runCmd(dir, "templ", "generate"); err != nil {
		t.Fatalf("templ generate failed: %v", err)
	}

	// Resolve go.mod dependencies and build.
	if err := runCmd(dir, "go", "mod", "tidy"); err != nil {
		t.Fatalf("go mod tidy failed: %v", err)
	}
	if err := runCmd(dir, "go", "build", "-o", os.DevNull, "."); err != nil {
		t.Fatalf("go build failed: %v", err)
	}
}

// TestPocketbaseTemplatesGenerateAndCompile_HTML verifies that the generated
// PocketBase backend code (html/template variant) compiles successfully.
func TestPocketbaseTemplatesGenerateAndCompile_HTML(t *testing.T) {
	// integration test - requires network

	dir := t.TempDir()
	generatePocketbaseBackend(t, dir, false)

	// Resolve go.mod dependencies and build.
	if err := runCmd(dir, "go", "mod", "tidy"); err != nil {
		t.Fatalf("go mod tidy failed: %v", err)
	}
	if err := runCmd(dir, "go", "build", "-o", os.DevNull, "."); err != nil {
		t.Fatalf("go build failed: %v", err)
	}
}

// TestPocketbaseTemplatesRender verifies that the gotmpl templates can be
// rendered without errors for both the templ and html/template variants.
func TestPocketbaseTemplatesRender(t *testing.T) {
	for _, useTempl := range []bool{true, false} {
		label := "html"
		if useTempl {
			label = "templ"
		}
		t.Run(label, func(t *testing.T) {
			dir := t.TempDir()
			generatePocketbaseBackend(t, dir, useTempl)

			// Verify key files were generated.
			assert.FileExists(t, filepath.Join(dir, "server.go"))
			assert.FileExists(t, filepath.Join(dir, "handlers.go"))
			assert.FileExists(t, filepath.Join(dir, "go.mod"))
		})
	}
}

// runCmd runs a command in the given directory.
func runCmd(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}