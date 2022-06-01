package cli

import (
	"fmt"
	"github.com/muesli/termenv"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

func fromGeneralToDetail(m GuiModelGeneral) {
	var selectedNews []entities.ArticlePreview

	for i, n := range m.news {
		// Is this choice selected?
		if _, ok := m.selected[i]; ok {
			selectedNews = append(selectedNews, n)
		}
	}

	p := tea.NewProgram(InitDetailModel(m.logger, selectedNews),
		tea.WithAltScreen()) // use the full size of the terminal in its "alternate screen buffer"
	//tea.WithMouseCellMotion()) // turn on mouse support so we can track the mouse wheel
	if err := p.Start(); err != nil {
		m.logger.Printf("error implementing cli mode: %v", err)
		os.Exit(1)
	}

	termenv.ClearScreen()
	fmt.Printf("\n%s", m.View())

}

func fromDetailToGeneral(m GuiModelDetail) {

}
