package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

var (
	styleTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "25", Dark: "86"})

	styleMenu = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "238", Dark: "250"})

	styleMenuSel = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "255", Dark: "230"}).
			Background(lipgloss.AdaptiveColor{Light: "62", Dark: "63"}).
			Padding(0, 1)

	styleErr = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "160", Dark: "203"})

	styleOK = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "28", Dark: "121"})

	styleHelp = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "241", Dark: "244"})
)

func applyTextareaTheme(t *textarea.Model) {
	// Catppuccin Latte / Mocha–inspired single-theme editor (editable JSON, no token highlights).
	fs := t.FocusedStyle
	fs.Base = lipgloss.NewStyle().
		Background(lipgloss.AdaptiveColor{Light: "#eff1f5", Dark: "#1e1e2e"})
	fs.Text = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#4c4f69", Dark: "#cdd6f4"})
	fs.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.AdaptiveColor{Light: "#dce0e8", Dark: "#313244"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#4c4f69", Dark: "#cdd6f4"})
	ln := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#8c8fa1", Dark: "#6c7086"})
	fs.LineNumber = ln
	fs.CursorLineNumber = ln
	fs.Placeholder = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#acb0be", Dark: "#9399b2"}).
		Italic(true)
	t.FocusedStyle = fs

	bs := t.BlurredStyle
	bs.Base = fs.Base
	bs.Text = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#5c5f77", Dark: "#a6adc8"})
	t.BlurredStyle = bs
}
