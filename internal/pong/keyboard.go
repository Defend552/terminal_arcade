package pong

import (
	listcomponent "terminal_arcade/internal/list_component"

	tea "github.com/charmbracelet/bubbletea"
)

var keyActions = map[string]func(*Model) tea.Cmd{
	"up":   func(m *Model) tea.Cmd { PaddleUp(&m.LeftPaddle); return nil },
	"k":    func(m *Model) tea.Cmd { PaddleUp(&m.LeftPaddle); return nil },
	"down": func(m *Model) tea.Cmd { PaddleDown(&m.LeftPaddle, &m.Board); return nil },
	"j":    func(m *Model) tea.Cmd { PaddleDown(&m.LeftPaddle, &m.Board); return nil },
	"q":    func(m *Model) tea.Cmd { return tea.Quit },
	"ctrl+c": func(m *Model) tea.Cmd {
		return tea.Quit
	},
	"enter": func(m *Model) tea.Cmd {
		if m.Difficulty == Select {
			SelectDifficulty(m)
			return nil
		}
		return nil
	},
}

func HandleKey(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	if action, ok := keyActions[msg.String()]; ok {
		cmd := action(&m)
		return m, cmd
	}

	return m, nil
}

func ExitApplication() tea.Cmd {
	return tea.Quit
}

func PaddleUp(p *Paddle) {
	if p.Y > 0 {
		p.Y--
	}
}

func PaddleDown(p *Paddle, board *GameBoard) {
	if p.Y < board.Height-4 {
		p.Y++
	}
}

func SelectDifficulty(m *Model) {
	selected := listcomponent.SelectedString(m.DifficultyList)

	switch selected {
	case "Easy":
		m.Difficulty = Easy
	case "Medium":
		m.Difficulty = Medium
	case "Hard":
		m.Difficulty = Hard
	}
}
