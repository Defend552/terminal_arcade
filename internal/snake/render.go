package snake

import (
	"fmt"
	"strings"
	"terminal_arcade/internal/utils"
)

func Render(m Model) string {
	var b strings.Builder

	b.WriteString("Terminal Snake\n\n")

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
			case gameX == m.Snake[0].X &&
				gameY == m.Snake[0].Y:
				b.WriteString(PrintHead(m.LastMove))
			case m.isBody(gameX, gameY):
				b.WriteString("o")
			case gameX == m.Food.X &&
				gameY == m.Food.Y:
				b.WriteString("X")
			case gameX == totalBoardWidth+2:
				sidebarLine := m.PrintGameState(y)
				fmt.Fprintf(&b, "%-*s", sideBarWidth, sidebarLine)
			default:
				b.WriteString(" ")
			}
		}
		b.WriteString("\n")
	}

	b.WriteString("Press CTRL-c or q to exit\n")

	return b.String()
}

func (m Model) PrintGameState(y int) string {
	switch y {
	case 1:
		return "Game State Values"
	case 2:
		return fmt.Sprintf("Snake: (%d,%d)", m.Snake[0].X, m.Snake[0].Y)
	case 3:
		return fmt.Sprintf("Last Move %s", DifficultyToString(m.LastMove))
	case 4:
		if m.SnakeTail != nil {
			return fmt.Sprintf("Snake Tail: (%d,%d)", m.SnakeTail.X, m.SnakeTail.Y)
		} else {
			return ""
		}
	default:
		return ""
	}
}

func (m Model) PrintGameOverScreen() string {
	var b strings.Builder
	if m.CheckIfGameIsWon() {
		b.WriteString(utils.RenderASCII("You Won!"))
	} else {
		b.WriteString(utils.RenderASCII("You Lost"))
	}
	b.WriteString("\n")
	b.WriteString("Press r to restart game\n")
	b.WriteString("\n")
	b.WriteString("Press CTRL-c or q to exit\n")

	return b.String()
}
