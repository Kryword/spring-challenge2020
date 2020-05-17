package tests

import (
	"../game"
	"testing"
)

func TestPositionCreation(t *testing.T) {
	var x, y = 1, 2
	p := new(game.Position)
	p.X, p.Y = x, y
	if p.X != 1 || p.Y != 2 {
		t.Errorf("Position creation [%d,%d] not equal to expected %+v", x, y, p)
	}
}

func TestManhattanDistanceOfPositions(t *testing.T) {
	pos1 := new(game.Position)
	pos2 := new(game.Position)
	pos1.X, pos1.Y = 1, 1
	pos2.X, pos2.Y = 3, 3
	expectedDist := game.Abs(pos1.X-pos2.X) + game.Abs(pos1.Y-pos2.Y)
	res := pos1.Dist(pos2)
	if res != expectedDist {
		t.Errorf("Manhattan distance %d not equal to expected %d", res, expectedDist)
	}
}
