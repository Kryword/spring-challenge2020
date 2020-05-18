package game

import (
	"sort"
	"strings"
)

type Cell struct {
	Pos   Position
	Type  CellType
	Taken bool
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
	var pacs []*Pac
	if newPac.Mine {
		pacs = m.MyPacs
	} else {
		pacs = m.EnemyPacs
	}
	for i := 0; i < len(pacs); i++ {
		pac := pacs[i]
		if pac.Id == newPac.Id {
			// Update previous pac and new pos cell to empty
			m.UpdateCell(Cell{
				Pos:  pac.Pos,
				Type: Empty,
			})
			m.UpdateCell(Cell{
				Pos:  newPac.Pos,
				Type: Empty,
			})
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
		if newPac.Mine {
			m.MyPacs = append(pacs, newPac)
		} else {
			m.EnemyPacs = append(pacs, newPac)
		}
	}
}

func (m *Map) GetCellsSortedByDist(position *Position, cellTypes ...CellType) []Cell {
	result := make([]Cell, 0, m.Width*m.Height)
	for i := 0; i < len(m.Grid); i++ {
		cell := m.Grid[i]
		if !cell.Taken {
			for _, cellType := range cellTypes {
				if cell.Type == cellType {
					result = append(result, cell)
				}
			}
		}
	}

	sort.Slice(result[:], func(i, j int) bool {
		return result[i].Pos.Dist(position) < result[j].Pos.Dist(position)
	})
	return result
}
