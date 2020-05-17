package tests

import (
	"../game"
	"testing"
)

func TestPacCreation(t *testing.T) {
	expX, expY := 1, 5
	expId := 1
	expCooldown, expSpeedTurns := 0, 0
	expMine := false
	expType := game.Rock

	pac := new(game.Pac)
	pac.Init(expId, expX, expY, expType, expCooldown, expSpeedTurns, expMine)
	if pac.Id != expId || pac.Pos.X != expX || pac.Pos.Y != expY || pac.PacType != expType || pac.Cooldown != expCooldown || pac.SpeedTurns != expSpeedTurns || pac.Mine != expMine {
		t.Errorf("Pac not created correctly %+v different to expected [%d %d %d %d %d %d %t]", pac, expId, expX, expY, expType, expCooldown, expSpeedTurns, expMine)
	}
}

func TestWinningLosingRPS(t *testing.T) {
	p := game.Rock
	expWin, expLose := game.Paper, game.Scissors
	resWin, resLose := p.GetWinningLosingType()
	if expWin != resWin && expLose != resLose {
		t.Errorf("Winning losing RPS not correct, got: [%d, %d], want: [%d, %d]", resWin, resLose, expWin, expLose)
	}

	p = game.Paper
	expWin, expLose = game.Scissors, game.Rock
	resWin, resLose = p.GetWinningLosingType()
	if expWin != resWin && expLose != resLose {
		t.Errorf("Winning losing RPS not correct, got: [%d, %d], want: [%d, %d]", resWin, resLose, expWin, expLose)
	}

	p = game.Scissors
	expWin, expLose = game.Rock, game.Paper
	resWin, resLose = p.GetWinningLosingType()
	if expWin != resWin && expLose != resLose {
		t.Errorf("Winning losing RPS not correct, got: [%d, %d], want: [%d, %d]", resWin, resLose, expWin, expLose)
	}
}

func TestGetMove(t *testing.T) {
	pac := new(game.Pac)
	pac.Init(0, 1, 1, game.Rock, 0, 0, true)
	pac.Target = game.Position{
		X: 3,
		Y: 1,
	}
	pac.CurrentMove = game.Move
	resMove := pac.GetMove()
	expMove := "MOVE 0 3 1 [3, 1]"
	if resMove != expMove{
		t.Errorf("Pac GetMove invalid, got: %s, want: %s", resMove, expMove)
	}

	pac.CurrentMove = game.Speed
	resMove = pac.GetMove()
	expMove = "SPEED 0"
	if resMove != expMove{
		t.Errorf("Pac GetMove invalid, got: %s, want: %s", resMove, expMove)
	}

	pac.CurrentMove = game.Switch
	pac.EnemyType = game.Scissors
	resMove = pac.GetMove()
	expMove = "SWITCH 0 ROCK"
	if resMove != expMove{
		t.Errorf("Pac GetMove invalid, got: %s, want: %s", resMove, expMove)
	}
}
