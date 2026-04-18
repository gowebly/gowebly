package helpers

import (
	"github.com/charmbracelet/lipgloss"
)

// StringStyle configures text styling options (color and bold).
type StringStyle struct {
	IsBold bool
	Color  lipgloss.AdaptiveColor
}

// FrameStyle configures frame styling options (padding and border color).
type FrameStyle struct {
	Padding []int
	Color   lipgloss.AdaptiveColor
}

// MakeStyled applies the given style to a string using lipgloss.
func MakeStyled(s string, style *StringStyle) string {
	if style == nil {
		return s
	}

	return lipgloss.NewStyle().Foreground(style.Color).Bold(style.IsBold).Render(s)
}

// MakeStyledFrame renders text with a left border frame using lipgloss.
func MakeStyledFrame(s string, style *FrameStyle) string {
	if style == nil {
		return s
	}

	return lipgloss.NewStyle().
		Padding(style.Padding...).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderLeftForeground(style.Color).
		Render(s)
}
