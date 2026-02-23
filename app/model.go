package app

import (
	"terminal_arcade/internal/menu"
	"terminal_arcade/internal/pong"
)

type Screen int

const (
	ScreenMenu Screen = iota
	ScreenPong
	ScreenTetris
	ScreenPacman
)

type Model struct {
	Screen Screen
	Menu   menu.Model
	Pong   pong.Model
}
