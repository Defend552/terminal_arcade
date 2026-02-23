package pong

import (
	"fmt"
	"strings"
	"terminal_arcade/internal/utils"
)

func View(m Model) string {
	var b strings.Builder

	b.WriteString("Terminal Pong\n\n")

	if m.CheckIfGameIsOver() {
		return m.PrintGameOverScreen()
	}

	sideBarWidth := 25
	totalBoardWidth := m.Board.Width + 2
	totalWidth := totalBoardWidth + sideBarWidth
	totalBoardHeight := m.Board.Height + 2

	for y := range totalBoardHeight {
		for x := range totalWidth {
			if y == 0 || y == totalBoardHeight-1 {
				if x == 0 || x == totalBoardWidth-1 {
					b.WriteString("+")
					continue
				} else {
					b.WriteString("-")
					continue
				}
			}
			if x == 0 || x == totalBoardWidth-1 || x == totalWidth-1 {
				b.WriteString("|")
				continue
			}

			gameX := x - 1
			gameY := y - 1

			switch {
			//makes the pong paddles
			case gameX == 2 &&
				gameY >= m.LeftPaddle.Y &&
				gameY < m.LeftPaddle.Y+m.LeftPaddle.Height:
				b.WriteString("|")

			case gameX == m.Board.Width-3 &&
				gameY >= m.RightPaddle.Y &&
				gameY < m.RightPaddle.Y+m.RightPaddle.Height:
				b.WriteString("|")

			//makes the ball
			case gameX == m.BallPosition.X &&
				gameY == m.BallPosition.Y:
				b.WriteString("O")

			//generates the game state
			case gameX == totalBoardWidth+2:
				sidebarLine := m.PrintGameState(y)
				fmt.Fprintf(&b, "%-*s", sideBarWidth, sidebarLine)
			default:
				b.WriteString(" ")
			}
		}
		b.WriteString("\n")
	}

	fmt.Fprintf(&b, "\nScore: %d - %d\n", m.Score.Player1, m.Score.Player2)
	b.WriteString("Press CTRL-c or q to exit\n")

	return b.String()
}

func (m Model) PrintGameState(y int) string {
	switch y {
	case 1:
		return "Game State Values"
	case 2:
		return fmt.Sprintf("Ball Position: (%d,%d)", m.BallPosition.X, m.BallPosition.Y)
	case 3:
		return fmt.Sprintf("Ball Velocity: (%d,%d)", m.BallVelocity.X, m.BallVelocity.Y)
	case 4:
		return fmt.Sprintf("Left Paddle Y: %d", m.LeftPaddle.Y)
	case 5:
		return fmt.Sprintf("Right Paddle Y: %d", m.RightPaddle.Y)
	default:
		return ""
	}
}

func (m Model) PrintGameOverScreen() string {
	var b strings.Builder
	b.WriteString(utils.RenderASCII("WINNER"))
	b.WriteString("\n")
	b.WriteString(utils.RenderASCII("SCORE"))
	b.WriteString("\n")
	b.WriteString(utils.RenderASCII(fmt.Sprintf("%d - %d", m.Score.Player1, m.Score.Player2)))
	b.WriteString("\n")
	b.WriteString("Press r to restart game\n")
	b.WriteString("\n")
	b.WriteString("Press CTRL-c or q to exit\n")

	return b.String()
}
