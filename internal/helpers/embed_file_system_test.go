package helpers

import (
	"embed"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

//go:embed *_test.go
var TestFiles embed.FS

//go:embed testdata/*
var TestTemplates embed.FS

func TestCopyFilesFromEmbedFS(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)

	require.NoError(t, CopyFilesFromEmbedFS(TestFiles, []EmbedFile{
		{
			EmbedFile:  "embed_file_system_test.go",
			OutputFile: "tmp/embed_file_system_test.go",
		},
	}))
	require.True(t, IsExistInFolder("tmp/embed_file_system_test.go", false))

	require.Error(t, CopyFilesFromEmbedFS(TestFiles, []EmbedFile{
		{
			EmbedFile:  "I_DO_NOT_EXIST.file",
			OutputFile: "tmp/I_DO_NOT_EXIST.file",
		},
	}))
	require.False(t, IsExistInFolder("tmp/I_DO_NOT_EXIST.file", false))

	_ = os.RemoveAll("tmp")
}

func TestGenerateFilesByTemplateFromEmbedFS(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)

	require.NoError(t, GenerateFilesByTemplateFromEmbedFS(TestTemplates, []EmbedTemplate{
		{
			EmbedFile:  "testdata/backend/go.mod.tmpl",
			OutputFile: "tmp/go.mod",
			Data: struct {
				ModuleName string
			}{
				"Testing",
			},
		},
		{
			EmbedFile:  "testdata/frontend/package.json.tmpl",
			OutputFile: "tmp/package.json",
			Data: struct {
				PackageName string
			}{
				"Testing",
			},
		},
		{
			EmbedFile:  "testdata/misc/gitignore.tmpl",
			OutputFile: "tmp/.gitignore",
			Data:       nil,
		},
	}))
	require.True(t, IsExistInFolder("tmp/go.mod", false))
	require.True(t, IsExistInFolder("tmp/package.json", false))
	require.True(t, IsExistInFolder("tmp/.gitignore", false))

	// Already exists
	require.Error(t, GenerateFilesByTemplateFromEmbedFS(TestTemplates, []EmbedTemplate{
		{
			EmbedFile:  "testdata/backend/go.mod.tmpl",
			OutputFile: "tmp/go.mod",
			Data: struct {
				ModuleName string
			}{
				"Testing",
			},
		},
	}))
	require.True(t, IsExistInFolder("tmp/go.mod", false))

	// Non-embedded file
	require.Error(t, GenerateFilesByTemplateFromEmbedFS(TestTemplates, []EmbedTemplate{
		{
			EmbedFile:  "testdata/backend/I_DO_NOT_EXIST.file",
			OutputFile: "tmp/I_DO_NOT_EXIST.file",
			Data:       nil,
		},
	}))
	require.False(t, IsExistInFolder("tmp/I_DO_NOT_EXIST.file", false))

	_ = os.RemoveAll("tmp")
}
