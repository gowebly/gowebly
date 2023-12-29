package helpers

import (
	"github.com/charmbracelet/lipgloss"
)

// StringStyle represents a styled string.
type StringStyle struct {
	IsBold bool
	Color  lipgloss.AdaptiveColor
}

// FrameStyle represents a styled frame.
type FrameStyle struct {
	Padding []int
	Color   lipgloss.AdaptiveColor
}

// MakeStyled returns a styled string.
func MakeStyled(s string, style *StringStyle) string {
	if style == nil {
		return s // no style provided, return the original string
	}

	return lipgloss.NewStyle().Foreground(style.Color).Bold(style.IsBold).Render(s)
}

func MakeStyledFrame(s string, style *FrameStyle) string {
	if style == nil {
		return s // no style provided, return the original string
	}

	return lipgloss.NewStyle().
		Padding(style.Padding...).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderLeftForeground(style.Color).
		Render(s)
}
