package pong

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

type Model struct {
	BallPosition Vector
	BallVelocity Vector
	RightPaddle  Paddle
	LeftPaddle   Paddle
	Board        GameBoard
	Score        ScoreBoard
}
