package snake

import (
	"math/rand/v2"

	tea "github.com/charmbracelet/bubbletea"
)

func StartGame() Model {
	return Model{
		Board:     GameBoard{Height: 24, Width: 40},
		Snake:     []Coordinate{{20, 12}},
		Food:      GenerateFood([]Coordinate{{20, 12}}, GameBoard{Height: 24, Width: 40}),
		SnakeTail: nil,
	}
}

func GenerateFood(s []Coordinate, b GameBoard) Coordinate {
	for {
		x := rand.IntN(b.Width)
		y := rand.IntN(b.Height)
		food := Coordinate{X: x, Y: y}

		if !isOnSnake(food, s) {
			return food
		}
	}
}

func isOnSnake(c Coordinate, snake []Coordinate) bool {
	for _, segment := range snake {
		if segment.X == c.X && segment.Y == c.Y {
			return true
		}
	}
	return false
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
