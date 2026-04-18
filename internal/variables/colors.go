package variables

import "github.com/charmbracelet/lipgloss"

// CLI theme colors using lipgloss adaptive colors for light/dark mode support.
var (
	ColorBlue   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	ColorGreen  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	ColorYellow = lipgloss.AdaptiveColor{Light: "#E0C900", Dark: "#FFE600"}
	ColorRed    = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	ColorGrey   = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#B6B4AD"}
)
