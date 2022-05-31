package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

var (
	grey         = lipgloss.AdaptiveColor{Light: "#111111", Dark: "#fafafa"}
	red          = lipgloss.Color("#f50202")
	highlight    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	articleWidth = 150
	cardWidth    = 40

	articleStyle = lipgloss.NewStyle().
			Border(border, true).
			BorderForeground(grey).
			Padding(1, 1)

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
		BorderForeground(red).
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

	titleArticleStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(red).
				Margin(1, 0, 1, 0).
				Height(1).
				Width(articleWidth)

	articleDataStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(grey).
				Margin(1, 0, 1, 0).
				Height(1).
				Width(articleWidth)

	textArticleStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(grey).
				Margin(1, 0, 1, 0).
				Height(15).
				Width(articleWidth)

	titleInnerStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(grey).
			Height(4).
			Width(cardWidth)

	subtitleInnerStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(lipgloss.Color("#aaaaaa")).
				Height(2).
				Width(cardWidth)
)

func MakeArticle(articlePreview entities.ArticlePreview) string {
	return articleStyle.Render(lipgloss.JoinVertical(lipgloss.Left, titleArticleStyle.Render(articlePreview.Title),
		articleDataStyle.Render(articlePreview.Subtitle),
		articleDataStyle.Render(articlePreview.Article.Date),
		articleDataStyle.Render(articlePreview.Article.Author),
		textArticleStyle.Render(articlePreview.Article.Text)))
}

func MakeCard(articlePreview entities.ArticlePreview, isActive bool) string {
	if isActive {
		return card.Render(lipgloss.JoinVertical(lipgloss.Left, titleInnerStyle.Render(articlePreview.Title), subtitleInnerStyle.Render(articlePreview.Subtitle)))
	}
	return cardActive.Render(lipgloss.JoinVertical(lipgloss.Left, titleInnerStyle.Render(articlePreview.Title), subtitleInnerStyle.Render(articlePreview.Subtitle)))
}

func MakeMenu(menuText string) string {
	return menu.Render(menuText)
}
