package helpers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// StyledOutput represents a struct for the styled output.
type StyledOutput struct {
	Text, State, Style string
}

// PrintStyledBlock prints block of the styled output of the given []string by
// the PrintStyled function.
func PrintStyledBlock(outputs []StyledOutput) {
	for _, o := range outputs {
		PrintStyled(o.Text, o.State, o.Style)
	}
}

// PrintStyled prints styled output for the given string by a lipgloss.Style
// template.
func PrintStyled(s, state, style string) {
	// Set lipgloss colors.
	successColor := lipgloss.AdaptiveColor{Light: "#16a34a", Dark: "#4ade80"}
	warningColor := lipgloss.AdaptiveColor{Light: "#ca8a04", Dark: "#facc15"}
	errorColor := lipgloss.AdaptiveColor{Light: "#dc2626", Dark: "#f87171"}
	infoColor := lipgloss.AdaptiveColor{Light: "#4b5563", Dark: "#9ca3af"}

	// Create a new lipgloss style instance.
	lg := lipgloss.NewStyle()

	// Switch between states.
	switch state {
	case "info":
		state = lg.Foreground(infoColor).SetString("– ").String()
	case "wait":
		state = lg.Foreground(warningColor).SetString("○ ").String()
	case "success":
		state = lg.Foreground(successColor).SetString("✓ ").String()
	case "error":
		state = lg.Foreground(errorColor).SetString("✕ ").String()
	case "warning":
		state = lg.Foreground(warningColor).SetString("‼ ").String()
	default:
		state = lg.SetString("").String()
	}

	// Concat state with the given string.
	concatStrings := state + s

	// Switch between styles.
	switch style {
	case "margin-top-bottom":
		s = renderOutput(concatStrings, lg.MarginTop(1).MarginBottom(1))
	case "margin-top":
		s = renderOutput(concatStrings, lg.MarginTop(1))
	case "margin-bottom":
		s = renderOutput(concatStrings, lg.MarginBottom(1))
	case "margin-left":
		s = renderOutput(concatStrings, lg.MarginLeft(1))
	case "margin-left-2":
		s = renderOutput(concatStrings, lg.MarginLeft(2))
	default:
		s = concatStrings
	}

	// Print styled output.
	fmt.Println(s)
}

// render renders a styled string with a given lipgloss.Style template
// using "charmbracelet/lipgloss" package.
func renderOutput(str string, template lipgloss.Style) string {
	return template.Render(str)
}
