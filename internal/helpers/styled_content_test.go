package helpers

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

func TestMakeStyled(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		style    *StringStyle
		expected string
	}{
		{
			name:     "Empty string with no style",
			input:    "",
			style:    nil,
			expected: "",
		},
		{
			name:     "Non-empty string with no style",
			input:    "Hello, world!",
			style:    nil,
			expected: "Hello, world!",
		},
		{
			name:  "Non-empty string with style",
			input: "Hello, world!",
			style: &StringStyle{
				Color:  lipgloss.AdaptiveColor{Light: "red", Dark: "red"},
				IsBold: true,
			},
			expected: "Hello, world!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MakeStyled(tc.input, tc.style)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMakeStyledFrame(t *testing.T) {
	// Test case 1: No style provided, return the original string
	result := MakeStyledFrame("Hello, World!", nil)
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 2: Style provided, return the styled string
	style := &FrameStyle{
		Padding: []int{0},
		Color:   lipgloss.AdaptiveColor{Light: "red", Dark: "red"},
	}
	result = MakeStyledFrame("Hello, World!", style)
	expected = "│Hello, World!"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
