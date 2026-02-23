package app

import (
	"terminal_arcade/menu"
	"terminal_arcade/pong"
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
