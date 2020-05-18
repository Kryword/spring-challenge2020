package game

import (
	"fmt"
)

type Pac struct {
	Id          int
	Pos         Position
	PacType     PacType
	Cooldown    int
	SpeedTurns  int
	Mine        bool
	Target      Position
	EnemyType   PacType
	CurrentMove PacMove
}

func (p *Pac) Init(id, x, y int, pacType PacType, cooldown, speedTurns int, mine bool) {
	p.Id = id
	p.Pos = Position{
		X: x,
		Y: y,
	}
	p.PacType = pacType
	p.Cooldown = cooldown
	p.SpeedTurns = speedTurns
	p.Mine = mine
}

func (p Pac) GetMove() string {
	switch p.CurrentMove {
	case Move:
		return fmt.Sprintf("MOVE %d %d %d [%d, %d]", p.Id, p.Target.X, p.Target.Y, p.Target.X, p.Target.Y)
	case Speed:
		return fmt.Sprintf("SPEED %d", p.Id)
	case Switch:
		winType, _ := p.EnemyType.GetWinningLosingType()
		return fmt.Sprintf("SWITCH %d %s", p.Id, winType.ToStr())
	default:
		return ""
	}
}
