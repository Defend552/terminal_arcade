package menu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func New() Model {
	items := []list.Item{
		item("Pong"),
		item("Tetris"),
		item("Pac-Man"),
	}

	l := list.New(items, itemDelegate{}, 20, 14)
	l.Title = "Select a Game"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return Model{List: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "enter" {
			if i, ok := m.List.SelectedItem().(item); ok {
				m.Choice = string(i)
			}
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return "\n" + m.List.View()
}
