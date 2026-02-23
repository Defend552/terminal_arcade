package utils

import "strings"

func RenderASCII(text string) string {
	var b strings.Builder

	height := 5

	for row := range height {
		for _, char := range strings.ToUpper(text) {
			if letter, ok := asciiFont[char]; ok {
				b.WriteString(letter[row])
				b.WriteString("  ")
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}
