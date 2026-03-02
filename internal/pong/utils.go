package pong

func DifficultyToString(d Difficulty) string {
	switch d {
	case Select:
		return "Select"
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	default:
		return "Unknown"
	}
}
