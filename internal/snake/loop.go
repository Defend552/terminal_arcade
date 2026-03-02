package snake

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func Tick(m Model) tea.Cmd {
	speed := time.Millisecond * 100
	return tea.Tick(speed, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) UpdateOnTick() (Model, tea.Cmd) {
	m = m.UpdateSnakePosition()
	return m, Tick(m)
}

func (m Model) UpdateSnakePosition() Model {
	if m.CheckIfGameIsOver() {
		return m
	}
	m = m.MoveSnakeLastMove()
	m = m.DoesSnakeGrow()
	return m
}

func (m Model) MoveSnakeLastMove() Model {
	snake := []Coordinate{}
	head := m.Snake[0]
	tail := m.Snake[len(m.Snake)-1]
	newHead := Coordinate{}

	switch m.LastMove {
	case Up:
		newHead = Coordinate{X: head.X, Y: head.Y - 1}
	case Down:
		newHead = Coordinate{X: head.X, Y: head.Y + 1}
	case Left:
		newHead = Coordinate{X: head.X - 1, Y: head.Y}
	case Right:
		newHead = Coordinate{X: head.X + 1, Y: head.Y}
	}

	snake = append([]Coordinate{newHead}, (m.Snake)...)
	snake = snake[:len(snake)-1]
	m.Snake = snake
	m.SnakeTail = &tail
	return m
}

func (m Model) DoesSnakeGrow() Model {
	snake_head := m.Snake[0]
	if m.Food.X == snake_head.X && m.Food.Y == snake_head.Y {
		m.Food = GenerateFood()
		m.Snake = append(m.Snake, *m.SnakeTail)
	}
	return m
}

// TODO: Need to update logic so it will check if the food can't be generated any more
func (m Model) CheckIfGameIsWon() bool {
	return false
}

func (m Model) CheckIfGameIsOver() bool {
	return m.Snake[0].X < 0 || m.Snake[0].X == m.Board.Width || m.Snake[0].Y < 0 || m.Snake[0].Y == m.Board.Height || m.CheckForCollisions()
}

func (m Model) CheckForCollisions() bool {
	head := m.Snake[0]

	for i := 1; i < len(m.Snake); i++ {
		if m.Snake[i] == head {
			return true
		}
	}
	return false
}

func (m Model) isBody(x, y int) bool {
	for i := 1; i < len(m.Snake); i++ {
		if m.Snake[i].X == x && m.Snake[i].Y == y {
			return true
		}
	}
	return false
}
