package pong

import "math/rand/v2"

// type Difficulty int
//
// const (
// 	Easy Difficulty = iota
// 	Medium
// 	Hard
// )

func (m Model) UpdateAIPaddle() Model {

	if checkForRandomness(m) {
		return m
	}

	paddleCenter := m.RightPaddle.Y + m.RightPaddle.Height/2

	if m.BallPosition.Y < paddleCenter {
		m.RightPaddle.Y--
	} else if m.BallPosition.Y > paddleCenter {
		m.RightPaddle.Y++
	}

	// clamp to board
	if m.RightPaddle.Y < 0 {
		m.RightPaddle.Y = 0
	}
	if m.RightPaddle.Y > m.Board.Height-m.RightPaddle.Height {
		m.RightPaddle.Y = m.Board.Height - m.RightPaddle.Height
	}

	return m
}

func predictBallY(m Model) int {
	x := m.BallPosition.X
	y := m.BallPosition.Y
	vx := m.BallVelocity.X
	vy := m.BallVelocity.Y

	for x < m.Board.Width-3 {
		x += vx
		y += vy

		if y <= 0 || y >= m.Board.Height-1 {
			vy *= -1
		}
	}

	return y
}

func checkForRandomness(m Model) bool {
	chances := 0.15
	if m.Difficulty == Easy {
		chances = 0.25
	}
	return rand.Float32() < float32(chances)
}

// targetY := predictBallY(m)
// paddleCenter := m.RightPaddle.Y + m.RightPaddle.Height/2
//
// if targetY < paddleCenter {
// 	m.RightPaddle.Y--
// } else if targetY > paddleCenter {
// 	m.RightPaddle.Y++
// }
