package cli

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

type GuiModelGeneral struct {
	logger   *log.Logger
	news     []entities.ArticlePreview // model behind rendered news
	cursor   int                       // which to-do list item our cursor is pointing at
	selected map[int]struct{}          // which to-do items are selected
}

func InitGeneralModel(l *log.Logger, news []entities.ArticlePreview) GuiModelGeneral {
	return GuiModelGeneral{
		logger: l,
		news:   news,
		// A map which indicates which articles are selected
		selected: make(map[int]struct{}),
	}
}

func (m GuiModelGeneral) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m GuiModelGeneral) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// Go to the detailed view of selected articles
		case "a":
			fromGeneralToDetail(m)
			return m, nil

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.news)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated GuiModelGeneral to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m GuiModelGeneral) View() string {
	// The header
	s := ""

	// Render the row
	//doc := strings.Builder{}
	rows := ""

	// Iterate over our newsGui
	for i, n := range m.news {

		// Is the cursor pointing at this choice?
		cursor := "" // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		if i != 0 && i%3 == 0 {
			s = lipgloss.JoinVertical(lipgloss.Left, s, rows)
			rows = ""
		}

		rows = lipgloss.JoinHorizontal(lipgloss.Top, rows, fmt.Sprintf("%d. %s [%s]\n%s ",
			i+1,
			cursor,
			checked,
			MakeCard(n, checked == "x")))
	}

	// menu
	s = lipgloss.JoinHorizontal(lipgloss.Top, MakeMenu(menuContentGeneral), lipgloss.JoinVertical(lipgloss.Left, s, rows))

	// Send the UI for rendering
	return s
}
