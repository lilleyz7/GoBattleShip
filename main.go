package main

import (
	types "BattleShip/Types"
	"fmt"
)

const length = 5
const width = 5

func PickShipLocations(p *types.Player, b *types.GameBoard) {
	var xVal int
	var yVal int
	var alignment string
	i := 0

	for i < 4 {
		fmt.Println("You will select 4 points on the board and select to face the ships vert or horizontally")
		fmt.Println("Select your first x value:")
		fmt.Scanln(&xVal)
		fmt.Println("Select your first y value:")
		fmt.Scanln(&yVal)
		fmt.Println("Vertical or horizontal? Enter v or h")
		fmt.Scanln(&alignment)
		shipSet := SetShips(xVal, yVal, alignment, b)
		if shipSet == false {
			fmt.Print("Not a legal coordinate for placement try again")
		} else {
			i++
		}
	}

}

func SetShips(x int, y int, a string, b *types.GameBoard) bool {
	placement := b.SetLocation(x, y)
	if placement == true {
		if a == "v" {
			b.SetLocation(x, y+1)
			b.SetLocation(x, y-1)
			return true
		} else {
			b.SetLocation(x+1, y)
			b.SetLocation(x-1, y)
			return true
		}
	} else {
		return false
	}

}

func RunGame(player *types.Player, ai *types.Player, pBoard *types.GameBoard, aBoard *types.GameBoard, state *types.Gamestate) {
	for {
		pBoard.Displayboard()
		break

	}
}

func main() {
	//state := NewGameState()
	player := types.NewPlayer("Zach", 12)
	//aiPlayer := NewPlayer("AI", 12)
	playerBoard := types.NewGameBoard(length, width)
	playerBoard.Displayboard()
	//aiBoard := NewGameBoard(width, length)
	//PickShipLocations(player, playerBoard)

	//RunGame(player, aiPlayer, playerBoard, aiBoard, state)
	fmt.Print(player.GetRemainingPieces())

}
