package game

import "fmt"
import "os"
import "bufio"

/**
 * Grab the pellets as fast as you can!
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: size of the grid
	// height: top left corner is (x=0, y=0)
	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)
	mapInput := ""
	for i := 0; i < height; i++ {
		scanner.Scan()
		mapInput += scanner.Text() + "\n"
		//row := scanner.Text() // one line of the grid: space " " is floor, pound "#" is wall
	}
	gameMap := new(Map)
	gameMap.Init(mapInput, width, height)
	for {
		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)
		// visiblePacCount: all your pacs and enemy pacs in sight
		var visiblePacCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePacCount)

		for i := 0; i < visiblePacCount; i++ {
			// pacId: pac number (unique within a team)
			// mine: true if this pac is yours
			// x: position in the grid
			// y: position in the grid
			// typeId: unused in wood leagues
			// speedTurnsLeft: unused in wood leagues
			// abilityCooldown: unused in wood leagues
			var pacId int
			var mine bool
			var _mine int
			var x, y int
			var typeId string
			var speedTurnsLeft, abilityCooldown int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &pacId, &_mine, &x, &y, &typeId, &speedTurnsLeft, &abilityCooldown)
			mine = _mine != 0
			pac := new(Pac)
			pac.Init(pacId, x, y, GetPacTypeFromStr(typeId), abilityCooldown, speedTurnsLeft, mine)
			gameMap.UpdatePac(pac)
		}
		// visiblePelletCount: all pellets in sight
		var visiblePelletCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePelletCount)

		cells := make([]Cell, visiblePelletCount)
		for i := 0; i < visiblePelletCount; i++ {
			// value: amount of points this pellet is worth
			var x, y, value int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &x, &y, &value)
			cellType := Empty
			if value == 1 {
				cellType = Pellet
			} else if value == 10 {
				cellType = SuperPellet
			}
			cells[i] = Cell{
				Pos: Position{
					X: x,
					Y: y,
				},
				Type: cellType,
			}
			gameMap.UpdateCell(cells[i])
		}
		fmt.Fprintln(os.Stderr, gameMap.GridToStr())

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		for i := 0; i < len(gameMap.MyPacs); i++ {
			pac := gameMap.MyPacs[i]

			// Use speed everytime it is possible
			if pac.Cooldown == 0 {
				pac.CurrentMove = Speed
				continue
			}

			cells := gameMap.GetCellsSortedByDist(&pac.Pos, SuperPellet, Pellet)
			if len(cells) > 0 {
				// Pick closest pellet and target it
				pac.Target = cells[0].Pos
				if pac.SpeedTurns > 0 {
					// If speeding, pick a cell with 2 dist or more if possible
					for j := 0; j < len(cells); j++ {
						if cells[j].Pos.Dist(&pac.Pos) >= 2 {
							pac.Target = cells[j].Pos
						}
					}
				}
				pac.CurrentMove = Move
				continue
			}
		}

		// Print output for every pac
		outStr := ""
		for i := 0; i < len(gameMap.MyPacs); i++ {
			pac := gameMap.MyPacs[i]
			outStr += pac.GetMove() + "|"
		}
		fmt.Println(outStr)
	}
}
