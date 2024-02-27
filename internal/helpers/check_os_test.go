package helpers

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test IsWindows function
func TestIsWindows(t *testing.T) {
	assert.Equal(t, runtime.GOOS == "windows", IsWindows())
}
