package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

var (
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

	card = lipgloss.NewStyle().
		Border(border, true).
		BorderForeground(lipgloss.Color("#f50202")).
		Padding(0, 1)

	cardActive = lipgloss.NewStyle().
			Border(border, true).
			BorderForeground(highlight).
			Padding(0, 1)

	titleInnerStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(lipgloss.Color("#fafafa")).
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
