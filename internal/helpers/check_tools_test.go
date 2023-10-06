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
	testCase1 := CheckTools([]Tool{{"go", "version"}})
	output1, _ := exec.Command("go", "version").Output()
	require.EqualValues(t, testCase1[0].Output, strings.Trim(string(output1), "\n"))

	testCase2 := CheckTools([]Tool{{"unknown", "version"}})
	require.EqualValues(t, testCase2[0].Output, fmt.Sprintf(constants.ErrorHelperToolNotInstalled, "unknown"))
}
