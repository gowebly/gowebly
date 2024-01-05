package helpers

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFilesFromEmbedFS(t *testing.T) {
	// Test case 1: Copying a single file
	fs := embed.FS{} // replace with actual embed.FS implementation
	files := []EmbedFile{
		{
			EmbedFile:  "file1.txt",
			OutputFile: "output/file1.txt",
		},
	}
	assert.Error(t, CopyFilesFromEmbedFS(fs, files))

	// Test case 2: Copying multiple files
	files = []EmbedFile{
		{
			EmbedFile:  "file1.txt",
			OutputFile: "output/file1.txt",
		},
		{
			EmbedFile:  "file2.txt",
			OutputFile: "output/file2.txt",
		},
		{
			EmbedFile:  "file3.txt",
			OutputFile: "output/file3.txt",
		},
	}
	assert.Error(t, CopyFilesFromEmbedFS(fs, files))

	// Test case 3: Error case - file not found
	files = []EmbedFile{
		{
			EmbedFile:  "nonexistent.txt",
			OutputFile: "output/nonexistent.txt",
		},
	}
	assert.Error(t, CopyFilesFromEmbedFS(fs, files))
}

func TestGenerateFilesByTemplateFromEmbedFS(t *testing.T) {
	// Test case 1: Generating a single file
	fs := embed.FS{} // replace with actual embed.FS implementation
	template := EmbedTemplate{
		EmbedFile:  "file1.txt",
		OutputFile: "output/file1.txt",
	}
	assert.Error(t, GenerateFilesByTemplateFromEmbedFS(fs, []EmbedTemplate{template}))

	// Test case 2: Generating multiple files
	template = EmbedTemplate{
		EmbedFile:  "file1.txt",
		OutputFile: "output/file1.txt",
	}
	assert.Error(t, GenerateFilesByTemplateFromEmbedFS(fs, []EmbedTemplate{template}))

	// Test case 3: Error case - file not found
	template = EmbedTemplate{
		EmbedFile:  "nonexistent.txt",
		OutputFile: "output/nonexistent.txt",
	}
	assert.Error(t, GenerateFilesByTemplateFromEmbedFS(fs, []EmbedTemplate{template}))

	// Test case 4: Generating multiple files
	template1 := EmbedTemplate{
		EmbedFile:  "file1.txt",
		OutputFile: "output/file1.txt",
	}
	template2 := EmbedTemplate{
		EmbedFile:  "file2.txt",
		OutputFile: "output/file2.txt",
	}
	templates := []EmbedTemplate{template1, template2}
	assert.Error(t, GenerateFilesByTemplateFromEmbedFS(fs, templates))
}
