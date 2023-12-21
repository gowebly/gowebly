package helpers

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestIsExistInFolder(t *testing.T) {
	path, _ := os.Getwd()
	folder := filepath.Base(path)

	require.True(t, IsExistInFolder(path, true))
	require.False(t, IsExistInFolder(folder, true))
}

func TestMakeFile(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)

	testFileName := "tmp/test_make_file.tmp"
	require.NoError(t, MakeFile(File{
		Name: testFileName,
		Data: nil,
	}))

	require.True(t, IsExistInFolder(testFileName, false))

	require.Error(t, MakeFile(File{
		Name: "tmp/test_make_file.tmp",
		Data: nil,
	}))

	_ = os.RemoveAll("tmp")
}

func TestMakeFiles(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)

	require.NoError(t, MakeFiles([]File{
		{
			Name: "tmp/test_make_file1.tmp",
			Data: nil,
		},
		{
			Name: "tmp/test_make_file2.tmp",
			Data: nil,
		},
	}))

	require.True(t, IsExistInFolder("tmp/test_make_file1.tmp", false))
	require.True(t, IsExistInFolder("tmp/test_make_file2.tmp", false))

	require.Error(t, MakeFiles([]File{
		{
			Name: "tmp/test_make_file3.tmp",
			Data: nil,
		},
		{
			Name: "tmp/test_make_file1.tmp",
			Data: nil,
		},
	}))

	_ = os.RemoveAll("tmp")
}

func TestMakeFolders(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)

	require.NoError(t, MakeFolders("tmp/tmp1", "tmp/tmp2"))

	require.True(t, IsExistInFolder("tmp/tmp1", true))
	require.True(t, IsExistInFolder("tmp/tmp2", true))

	require.Error(t, MakeFolders("tmp/tmp3", "tmp/tmp1"))

	_ = os.RemoveAll("tmp")
}

func TestRemoveFiles(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)
	_ = MakeFiles([]File{
		{
			Name: "tmp/test_make_file1.tmp",
			Data: nil,
		},
		{
			Name: "tmp/test_make_file2.tmp",
			Data: nil,
		},
		{
			Name: "tmp/test_make_file3.tmp",
			Data: nil,
		},
	})

	require.NoError(t, RemoveFiles("tmp/test_make_file1.tmp", "tmp/test_make_file2.tmp"))

	require.False(t, IsExistInFolder("tmp/test_make_file1.tmp", false))
	require.False(t, IsExistInFolder("tmp/test_make_file2.tmp", false))
	require.True(t, IsExistInFolder("tmp/test_make_file3.tmp", false))

	require.Error(t, RemoveFiles("tmp/test_make_file3.tmp", "tmp/test_make_file4.tmp"))

	_ = os.RemoveAll("tmp")
}
