package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

var (
	grey        = lipgloss.AdaptiveColor{Light: "#111111", Dark: "#fafafa"}
	highlight   = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	columnWidth = 40

	border = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	borderRight = lipgloss.Border{
		Right: "│",
	}

	card = lipgloss.NewStyle().
		Border(border, true).
		BorderForeground(lipgloss.Color("#f50202")).
		Padding(0, 1)

	cardActive = lipgloss.NewStyle().
			Border(border, true).
			BorderForeground(highlight).
			Padding(0, 1)

	menu = lipgloss.NewStyle().
		Border(borderRight, true).
		BorderForeground(grey).
		Margin(0, 2).
		Padding(0, 2).
		Height(4 * 14)

	titleInnerStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(grey).
			Margin(1, 3, 0, 0).
			Padding(1, 2).
			Height(4).
			Width(columnWidth)

	subtitleInnerStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(lipgloss.Color("#aaaaaa")).
				Margin(1, 3, 0, 0).
				Padding(1, 2).
				Height(2).
				Width(columnWidth)
)

func MakeCard(articlePreview entities.ArticlePreview, isActive bool) string {
	if isActive {
		return card.Render(lipgloss.JoinVertical(lipgloss.Left, titleInnerStyle.Render(articlePreview.Title), subtitleInnerStyle.Render(articlePreview.Subtitle)))
	}
	return cardActive.Render(lipgloss.JoinVertical(lipgloss.Left, titleInnerStyle.Render(articlePreview.Title), subtitleInnerStyle.Render(articlePreview.Subtitle)))
}

func MakeMenu(menuText string) string {
	return menu.Render(menuText)
}
