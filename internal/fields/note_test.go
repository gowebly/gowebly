package fields

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWelcomeNote(t *testing.T) {
	require.NoError(t, WelcomeNote().Error())
}
