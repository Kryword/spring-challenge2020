package game

type PacType int
type PacMove int
type CellType int

const (
	Rock PacType = iota
	Paper
	Scissors
)

func (p PacType) GetWinningLosingType() (PacType, PacType) {
	switch p {
	case Rock:
		return Paper, Scissors
	case Paper:
		return Scissors, Rock
	case Scissors:
		return Rock, Paper
	// It should never go to default
	default:
		return Rock, Rock
	}
}

func (p PacType) ToStr() string{
	switch p {
	case Rock: return "ROCK"
	case Paper: return "PAPER"
	case Scissors: return "SCISSORS"
	default: return ""
	}
}

const (
	Move PacMove = iota
	Speed
	Switch
)


const (
	Pellet CellType = iota
	SuperPellet
	Empty
	Wall
	Unknown
)

func (c CellType) ToStr() string{
	switch c {
	case Pellet:
		return "."
	case SuperPellet:
		return "O"
	case Empty:
		return " "
	case Wall:
		return "#"
	case Unknown:
		return "?"
	// It should never be default
	default:
		return ""
	}
}
