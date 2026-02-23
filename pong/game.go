package pong

import (
	tea "github.com/charmbracelet/bubbletea"
)

func StartGame() Model {
	return Model{
		BallPosition: Vector{X: 40, Y: 12},
		BallVelocity: Vector{X: 1, Y: -1},
		RightPaddle:  Paddle{Y: 20, Height: 4},
		LeftPaddle:   Paddle{Y: 20, Height: 4},
		Board:        GameBoard{Height: 24, Width: 80},
		Score:        ScoreBoard{Player1: 0, Player2: 0},
	}
}

func (m Model) Init() tea.Cmd {
	return Tick()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return HandleKey(m, msg)

	case tickMsg:
		return m.UpdateOnTick()
	}

	return m, nil
}

func (m Model) View() string {
	return View(m)
}
