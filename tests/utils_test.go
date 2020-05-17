package tests

import "testing"
import "../game"

func TestAbsOfInt(t *testing.T) {
	res := game.Abs(-5)
	exp := 5
	if res != exp {
		t.Errorf("Result %d distinct to expected value %d", res, exp)
	}

	res = game.Abs(10)
	exp = 10
	if res != exp {
		t.Errorf("Result %d distinct to expected value %d", res, exp)
	}
}
