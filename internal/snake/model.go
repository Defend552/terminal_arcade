package snake

type GameBoard struct {
	Height int
	Width  int
}

type Coordinate struct {
	X int
	Y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Model struct {
	Board     GameBoard
	Snake     []Coordinate
	LastMove  Direction
	Food      Coordinate
	SnakeTail *Coordinate
}
