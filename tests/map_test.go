package tests

import (
	"../game"
	"testing"
)

func TestInitializeEmptyMap(t *testing.T) {
	mapInput := "## \n"
	width, height := 3, 1
	gameMap := new(game.Map)
	gameMap.Init(mapInput, width, height)

	if gameMap.Width != width || gameMap.Height != height{
		t.Errorf("Initialized map width, height incorrect, got: [%d,%d], want: [%d,%d]", gameMap.Width, gameMap.Height, width, height)
	}

	resGrid := gameMap.GridToStr()
	if resGrid != mapInput{
		t.Errorf("Map not initialized correctly, got: %s, want %s", resGrid, mapInput)
	}

	if gameMap.Grid[1].Type != game.Wall{
		t.Errorf("Map cells not initialized correctly, got: %+v, want: Wall type", gameMap.Grid[1])
	}
}

func TestUpdateCellInMap(t *testing.T) {
	mapInput :=
		"# #  \n" +
		"# # #\n" +
		"#   #\n"
	width, height := 5, 3
	gameMap := new(game.Map)
	gameMap.Init(mapInput, width, height)

	updatedCell := game.Cell{
		Pos:  game.Position{
			X: 3,
			Y: 1,
		},
		Type: game.SuperPellet,
	}
	gameMap.UpdateCell(updatedCell)
	resultCell := gameMap.GetCell(game.Position{
		X: 3,
		Y: 1,
	})

	if updatedCell != resultCell {
		t.Errorf("Updating cells fail, got: %+v, want: %+v", resultCell, updatedCell)
	}

	// O stands for SuperPellet
	expectedGrid :=
		"# #  \n" +
		"# #O#\n" +
		"#   #\n"
	resultingGrid := gameMap.GridToStr()
	if expectedGrid != resultingGrid{
		t.Errorf("Updating cells not updating grid, got: %s, want: %s", resultingGrid, expectedGrid)
	}
}

func TestUpdateSeveralCells(t *testing.T) {
	mapInput :=
		"# #  \n" +
		"#   #\n" +
		"# # #\n"
	width, height := 5, 3
	gameMap := new(game.Map)
	gameMap.Init(mapInput, width, height)
	cells := [...]game.Cell{game.Cell{
		Pos:  game.Position{
			X: 1,
			Y: 0,
		},
		Type: game.Pellet,
	}, game.Cell{
		Pos:  game.Position{
			X: 3,
			Y: 0,
		},
		Type: game.SuperPellet,
	}, game.Cell{
		Pos:  game.Position{
			X: 2,
			Y: 1,
		},
		Type: game.Unknown,
	}}
	for i:= 0; i < len(cells); i++{
		gameMap.UpdateCell(cells[i])
	}
	expectedGrid :=
			"#.#O \n" +
			"# ? #\n" +
			"# # #\n"
	resultGrid := gameMap.GridToStr()

	if resultGrid != expectedGrid{
		t.Errorf("Fail updating multiple cells in map, got: %s, want: %s", resultGrid, expectedGrid)
	}
}
