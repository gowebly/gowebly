package helpers

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gowebly/gowebly/internal/constants"
)

func TestCheckTools(t *testing.T) {
	output1, _ := exec.Command("go", "version").Output()

	require.EqualValues(
		t,
		CheckTools([]Tool{{"go", "version"}})[0].Output,
		strings.Trim(string(output1), "\n"),
	)
	require.EqualValues(
		t,
		CheckTools([]Tool{{"unknown", "version"}})[0].Output,
		fmt.Sprintf(constants.ErrorHelperToolNotInstalled, "unknown"),
	)
}
