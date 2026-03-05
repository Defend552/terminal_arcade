package snake

func PrintHead(d Direction) string {
	switch d {
	case Up:
		return "^"
	case Down:
		return "v"
	case Right:
		return ">"
	case Left:
		return "<"
	default:
		return "O"
	}
}

func DifficultyToString(d Direction) string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Right:
		return "Right"
	case Left:
		return "Left"
	default:
		return "None"
	}
}
