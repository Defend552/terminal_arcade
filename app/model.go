package app

import (
	"terminal_arcade/internal/menu"
	"terminal_arcade/internal/pong"
	"terminal_arcade/internal/snake"
)

type Screen int

const (
	ScreenMenu Screen = iota
	ScreenPong
	ScreenSnake
)

type Model struct {
	Screen Screen
	Menu   menu.Model
	Pong   pong.Model
	Snake  snake.Model
}
