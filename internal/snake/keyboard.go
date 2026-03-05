package snake

import (
	tea "github.com/charmbracelet/bubbletea"
)

var keyActions = map[string]func(*Model) tea.Cmd{
	"up": func(m *Model) tea.Cmd {
		SnakeUp(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"w": func(m *Model) tea.Cmd {
		SnakeUp(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"down": func(m *Model) tea.Cmd {
		SnakeDown(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"s": func(m *Model) tea.Cmd {
		SnakeDown(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"left": func(m *Model) tea.Cmd {
		SnakeLeft(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"a": func(m *Model) tea.Cmd {
		SnakeLeft(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"right": func(m *Model) tea.Cmd {
		SnakeRight(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"d": func(m *Model) tea.Cmd {
		SnakeRight(&m.Snake, &m.Board, &m.LastMove, &m.Food, m.SnakeTail, m.CheckIfGameIsOver())
		return nil
	},
	"q": func(m *Model) tea.Cmd { return tea.Quit },
	"ctrl+c": func(m *Model) tea.Cmd {
		return tea.Quit
	},
}

func HandleKey(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	if action, ok := keyActions[msg.String()]; ok {
		cmd := action(&m)
		return m, cmd
	}
	return m, nil
}

func SnakeUp(snake *[]Coordinate, board *GameBoard, direction *Direction, food, snakeTail *Coordinate, gameOver bool) {
	if !gameOver {
		*direction = Up
		head := (*snake)[0]
		snakeTail = &(*snake)[len(*snake)-1]
		newHead := Coordinate{X: head.X, Y: head.Y - 1}
		*snake = append([]Coordinate{newHead}, (*snake)...)
		*snake = (*snake)[:len(*snake)-1]
		if head.X == food.X && head.Y == food.Y {
			*food = GenerateFood(*snake, *board)
		}
	}
}

func SnakeDown(snake *[]Coordinate, board *GameBoard, direction *Direction, food, snakeTail *Coordinate, gameOver bool) {
	if !gameOver {
		*direction = Down
		head := (*snake)[0]
		snakeTail = &(*snake)[len(*snake)-1]
		newHead := Coordinate{X: head.X, Y: head.Y + 1}
		*snake = append([]Coordinate{newHead}, (*snake)...)
		*snake = (*snake)[:len(*snake)-1]
		if head.X == food.X && head.Y == food.Y {
			*food = GenerateFood(*snake, *board)
		}
	}
}

func SnakeLeft(snake *[]Coordinate, board *GameBoard, direction *Direction, food, snakeTail *Coordinate, gameOver bool) {
	if !gameOver {
		*direction = Left
		head := (*snake)[0]
		snakeTail = &(*snake)[len(*snake)-1]
		newHead := Coordinate{X: head.X - 1, Y: head.Y}
		*snake = append([]Coordinate{newHead}, (*snake)...)
		*snake = (*snake)[:len(*snake)-1]
		if head.X == food.X && head.Y == food.Y {
			*food = GenerateFood(*snake, *board)
		}
	}
}

func SnakeRight(snake *[]Coordinate, board *GameBoard, direction *Direction, food, snakeTail *Coordinate, gameOver bool) {
	if !gameOver {
		*direction = Right
		head := (*snake)[0]
		snakeTail = &(*snake)[len(*snake)-1]
		newHead := Coordinate{X: head.X + 1, Y: head.Y}
		*snake = append([]Coordinate{newHead}, (*snake)...)
		*snake = (*snake)[:len(*snake)-1]
		if head.X == food.X && head.Y == food.Y {
			*food = GenerateFood(*snake, *board)
		}
	}
}
