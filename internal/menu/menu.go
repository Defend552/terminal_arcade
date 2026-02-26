package menu

import (
	listcomponent "terminal_arcade/internal/list_component"

	tea "github.com/charmbracelet/bubbletea"
)

func New() Model {
	l := listcomponent.New(
		"Select a Game",
		[]string{
			"Pong",
			"Tetris",
			"Pac-Man",
		},
		20,
	)

	return Model{List: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "enter" {
			if i, ok := m.List.SelectedItem().(listcomponent.Item); ok {
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
