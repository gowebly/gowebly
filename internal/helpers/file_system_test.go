package helpers

import (
	"fmt"
	"os"
	"testing"

	"github.com/gowebly/gowebly/v2/internal/messages"
	"github.com/stretchr/testify/assert"
)

func TestIsExistInFolder(t *testing.T) {
	setup := func() {
		// Setup - create 'existing_folder' and 'existing_file.txt'
		os.Mkdir("existing_folder", 0o755)
		os.WriteFile("existing_file.txt", []byte("content"), 0o644)
	}
	tearDown := func() {
		// Teardown - remove 'existing_folder' and 'existing_file.txt' after test
		os.RemoveAll("existing_folder")
		os.Remove("existing_file.txt")
	}

	setup()
	defer tearDown()

	// Test case 1: Testing for an existing folder
	existingFolder := "existing_folder"
	if !IsExistInFolder(existingFolder, true) {
		t.Errorf("%s should exist in the current folder", existingFolder)
	}

	// Test case 2: Testing for a non-existing folder
	nonExistingFolder := "non_existing_folder"
	if IsExistInFolder(nonExistingFolder, true) {
		t.Errorf("%s should not exist in the current folder", nonExistingFolder)
	}

	// Test case 3: Testing for an existing file
	existingFile := "existing_file.txt"
	if !IsExistInFolder(existingFile, false) {
		t.Errorf("%s should exist in the current folder", existingFile)
	}

	// Test case 4: Testing for a non-existing file
	nonExistingFile := "non_existing_file.txt"
	if IsExistInFolder(nonExistingFile, false) {
		t.Errorf("%s should not exist in the current folder", nonExistingFile)
	}

	// Clean up.
	os.RemoveAll("existing_file.txt")
}

func TestMakeFile(t *testing.T) {
	// Test case 1: File doesn't exist, should create file successfully
	file1 := File{
		Name: "test.txt",
		Data: []byte("This is a test file"),
	}
	err := MakeFile(file1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test case 2: File already exists, should return an error
	file2 := File{
		Name: "test.txt",
		Data: []byte("This is another test file"),
	}
	err = MakeFile(file2)
	expectedErr := fmt.Errorf(messages.ErrorOSFileIsExists, file2.Name)
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}

	// Clean up.
	os.RemoveAll("test.txt")
}

func TestMakeFiles(t *testing.T) {
	// Test case 1: Empty files list
	files := []File{}
	err := MakeFiles(files)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Single file
	file := File{Name: "file1.txt", Data: []byte("Hello, World!")}
	files = []File{file}
	err = MakeFiles(files)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Clean up.
	os.RemoveAll("file1.txt")

	// Test case 3: Multiple files
	file1 := File{Name: "file1.txt", Data: []byte("Hello, World!")}
	file2 := File{Name: "file2.txt", Data: []byte("This is another file.")}
	files = []File{file1, file2}
	err = MakeFiles(files)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Clean up.
	os.RemoveAll("file1.txt")
	os.RemoveAll("file2.txt")
}

func TestMakeFolders(t *testing.T) {
	// Test case 1: Test with no folders to create
	err := MakeFolders()
	assert.NoError(t, err)

	// Test case 2: Test with single folder to create
	err = MakeFolders("folder1")
	assert.NoError(t, err)

	// Test case 3: Test with multiple folders to create
	err = MakeFolders("folder2", "folder3", "folder4")
	assert.NoError(t, err)

	// Clean up.
	os.RemoveAll("folder1")
	os.RemoveAll("folder2")
	os.RemoveAll("folder3")
	os.RemoveAll("folder4")
}
