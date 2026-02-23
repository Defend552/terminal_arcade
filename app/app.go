package app

import (
	"terminal_arcade/internal/menu"
	"terminal_arcade/internal/pong"

	tea "github.com/charmbracelet/bubbletea"
)

func StartApp() Model {
	return Model{
		Screen: ScreenMenu,
		Menu:   menu.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Menu.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.Screen {
	case ScreenMenu:
		var cmd tea.Cmd
		m.Menu, cmd = m.Menu.Update(msg)

		if m.Menu.Choice == "Pong" {
			m.Screen = ScreenPong
			m.Pong = pong.StartGame()
			return m, m.Pong.Init()
		}

		return m, cmd

	case ScreenPong:
		var cmd tea.Cmd
		m.Pong, cmd = m.Pong.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	switch m.Screen {
	case ScreenMenu:
		return m.Menu.View()
	case ScreenPong:
		return m.Pong.View()
	}

	return ""
}
