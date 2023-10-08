package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadFiles(t *testing.T) {
	_ = os.Mkdir("tmp", 0o755)

	require.NoError(
		t,
		DownloadFiles([]Download{
			{"https://go.dev/images/favicon-gopher.svg", "tmp/file.svg"},
		}),
	)

	_, err := os.Stat("tmp/file.svg")
	require.NoError(t, err)

	require.Error(
		t,
		DownloadFiles([]Download{
			{"https://go.dev/images/unknown.file", "tmp/unknown.file"},
		}),
	)
	require.Error(
		t,
		DownloadFiles([]Download{
			{"", ""},
		}),
	)

	_ = os.RemoveAll("tmp")
}
