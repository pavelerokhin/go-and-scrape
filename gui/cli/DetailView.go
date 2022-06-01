package cli

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"log"
)

type GuiModelDetail struct {
	logger   *log.Logger
	viewport viewport.Model
}

func InitDetailModel(l *log.Logger, news []entities.ArticlePreview) GuiModelDetail {
	// Render the articles
	articles := ""

	// Iterate over news
	for _, n := range news {
		// Is this choice selected?
		articles = lipgloss.JoinVertical(lipgloss.Top, articles,
			MakeArticle(n))
	}

	vp := viewport.New(10, 50)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))

	vp.SetContent(articles)

	return GuiModelDetail{
		logger:   l,
		viewport: vp,
	}
}

func (m GuiModelDetail) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m GuiModelDetail) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		return m, nil
	case tea.KeyMsg:
		// What was the actual key pressed:
		switch msg.String() {
		// These keys should exit the program.
		case "a":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func (m GuiModelDetail) View() string {
	vpi := m.viewport.View()
	return lipgloss.JoinHorizontal(lipgloss.Left, MakeMenu(menuContentDetail), vpi)
}
