package menu

import "github.com/charmbracelet/bubbles/list"

type Model struct {
	List     list.Model
	Choice   string
	Quitting bool
}
