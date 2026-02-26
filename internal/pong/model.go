package pong

import "github.com/charmbracelet/bubbles/list"

type Vector struct {
	X int
	Y int
}

type Paddle struct {
	Y      int
	Height int
}

type GameBoard struct {
	Height int
	Width  int
}

type ScoreBoard struct {
	Player1 int
	Player2 int
}

type Difficulty int

const (
	Select Difficulty = iota
	Easy
	Medium
	Hard
)

type Model struct {
	BallPosition   Vector
	BallVelocity   Vector
	RightPaddle    Paddle
	LeftPaddle     Paddle
	Board          GameBoard
	Score          ScoreBoard
	ScoreToWin     int
	DifficultyList list.Model
	Difficulty     Difficulty
}
