package snake

import (
	"math/rand/v2"

	tea "github.com/charmbracelet/bubbletea"
)

func StartGame() Model {
	return Model{
		Board:     GameBoard{Height: 24, Width: 40},
		Snake:     []Coordinate{{20, 12}},
		Food:      GenerateFood(),
		SnakeTail: nil,
	}
}

// TODO: Need to update logic so it will not be impacted by the snake
func GenerateFood() Coordinate {
	x := rand.IntN(40)
	y := rand.IntN(24)

	if x != 40 && y != 12 {
		return Coordinate{x, y}
	} else {
		GenerateFood()
	}
	return Coordinate{}
}

func (m Model) Init() tea.Cmd {
	return Tick(m)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			return StartGame(), nil
		default:
			return HandleKey(m, msg)
		}
	case tickMsg:
		return m.UpdateOnTick()
	}

	return m, nil
}

func (m Model) View() string {
	return Render(m)
}
