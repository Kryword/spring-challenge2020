package game

type Position struct {
	X, Y int
}

func (p Position) Dist(pos *Position) int {
	return Abs(p.X-pos.X) + Abs(p.Y-pos.Y)
}
