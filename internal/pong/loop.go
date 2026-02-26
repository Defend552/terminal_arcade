package pong

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func Tick(m Model) tea.Cmd {
	speed := time.Millisecond * 50
	if m.Difficulty == Hard {
		speed = time.Millisecond * 25
	}
	return tea.Tick(speed, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) UpdateOnTick() (Model, tea.Cmd) {
	m = m.UpdateBallPosition()
	return m, Tick(m)
}

func (m Model) UpdateBallPosition() Model {
	if m.CheckIfGameIsOver() {
		return m
	}

	m = m.MoveBall()
	m = m.TopBottomCollision()

	if m.CheckIfLeftPaddleIsHit() {
		m.BallVelocity.X = 1
		m.BallPosition.X = 3
	}

	if m.CheckIfRightPaddleIsHit() {
		m.BallVelocity.X = -1
		m.BallPosition.X = m.Board.Width - 4
	}

	if m.BallPosition.X > m.Board.Width {
		m.Score.Player1++
		m = m.ResetBall()
		return m
	}
	if m.BallPosition.X < 0 {
		m.Score.Player2++
		m = m.ResetBall()
		return m
	}
	if m.BallVelocity.X > 0 {
		m = m.UpdateAIPaddle()
	}

	return m
}

func (m Model) MoveBall() Model {
	m.BallPosition.Y += m.BallVelocity.Y
	m.BallPosition.X += m.BallVelocity.X
	return m
}

func (m Model) TopBottomCollision() Model {
	if m.BallPosition.Y >= m.Board.Height-1 || m.BallPosition.Y <= 0 {
		m.BallVelocity.Y *= -1
	}
	return m
}

func (m Model) ResetBall() Model {
	m.BallPosition.X = m.Board.Width / 2
	m.BallPosition.Y = m.Board.Height / 2
	return m
}

func (m Model) CheckIfLeftPaddleIsHit() bool {
	return m.BallVelocity.X < 0 &&
		m.BallPosition.X <= 2 &&
		m.BallPosition.Y >= m.LeftPaddle.Y &&
		m.BallPosition.Y < m.LeftPaddle.Y+m.LeftPaddle.Height
}

func (m Model) CheckIfRightPaddleIsHit() bool {
	return m.BallVelocity.X > 0 &&
		m.BallPosition.X >= m.Board.Width-3 &&
		m.BallPosition.Y >= m.RightPaddle.Y &&
		m.BallPosition.Y < m.RightPaddle.Y+m.RightPaddle.Height
}

func (m Model) CheckIfGameIsOver() bool {
	return m.Score.Player1 == m.ScoreToWin || m.Score.Player2 == m.ScoreToWin || m.Difficulty == Select
}
