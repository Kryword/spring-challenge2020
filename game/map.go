package game

import "strings"

type Cell struct {
	Pos  Position
	Type CellType
}

type Map struct {
	Grid                  []Cell
	Height, Width         int
	MyPoints, EnemyPoints int
	MyPacs                []*Pac
	EnemyPacs             []*Pac
}

func (m *Map) Init(nonParsedGrid string, width, height int) {
	m.Width = width
	m.Height = height
	// initialize layers for grid
	m.Grid = make([]Cell, height*width)

	rows := strings.Split(nonParsedGrid, "\n")
	if len(rows) > 0 {
		for i := 0; i < len(rows); i++ {
			chars := strings.Split(rows[i], "")
			for j := 0; j < len(chars); j++ {
				pos := Position{
					X: j,
					Y: i,
				}
				if chars[j] == "#" {
					m.Grid[i*width+j] = Cell{
						Pos:  pos,
						Type: Wall,
					}
				} else if chars[j] == " " {
					m.Grid[i*width+j] = Cell{
						Pos:  pos,
						Type: Empty,
					}
				}
			}
		}
	}
}

func (m Map) GridToStr() string {
	result := ""
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			result += m.Grid[i*m.Width+j].Type.ToStr()
		}
		result += "\n"
	}
	return result
}

func (m *Map) UpdateCell(newCell Cell) {
	x, y := newCell.Pos.X, newCell.Pos.Y
	m.Grid[y*m.Width+x] = newCell
}

func (m Map) GetCell(position Position) Cell {
	x, y := position.X, position.Y
	return m.Grid[y*m.Width+x]
}

func (m *Map) UpdatePac(newPac *Pac) {
	found := false
	if newPac.Mine {
		for i := 0; i < len(m.MyPacs); i++ {
			pac := m.MyPacs[i]
			if pac.Id == newPac.Id {
				// Update pac info
				pac.Pos = newPac.Pos
				pac.PacType = newPac.PacType
				pac.SpeedTurns = newPac.SpeedTurns
				pac.Cooldown = newPac.Cooldown
				// Target and EnemyType is not updating here, because that is done in other place
				found = true
				break
			}
		}
		if !found {
			m.MyPacs = append(m.MyPacs, newPac)
		}
	} else {
		// Update enemy pac
		for i := 0; i < len(m.EnemyPacs); i++ {
			pac := m.EnemyPacs[i]
			if pac.Id == newPac.Id {
				// Update pac info
				pac.Pos = newPac.Pos
				pac.PacType = newPac.PacType
				pac.SpeedTurns = newPac.SpeedTurns
				pac.Cooldown = newPac.Cooldown
				// Target and EnemyType is not updating here, because that is done in other place
				found = true
				break
			}
		}
		if !found {
			m.EnemyPacs = append(m.EnemyPacs, newPac)
		}
	}
}
