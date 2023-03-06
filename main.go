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

	for i < 3 {
		fmt.Println("You will select 4 points on the board and select to face the ships vert or horizontally")
		fmt.Println(i)
		fmt.Println("Select your first x value:")
		fmt.Scanln(&xVal)
		fmt.Println("Select your first y value:")
		fmt.Scanln(&yVal)
		fmt.Println("Vertical or horizontal? Enter v or h")
		fmt.Scanln(&alignment)
		position := types.NewPosition(xVal, yVal)
		shipSet := SetShips(position, alignment, b)
		if shipSet == false {
			fmt.Print("Not a legal coordinate for placement try again")
		} else {
			i++
		}
	}

}

func SetShips(p *types.Position, a string, b *types.GameBoard) bool {
	placement := b.SetShipLocations(p)
	if placement == true {
		if a == "v" {
			pos1, pos2 := p.UpdateVert()
			b.SetShipLocations(pos1)
			b.SetShipLocations(pos2)
			return true
		} else {
			pos1, pos2 := p.UpdateHoriz()
			b.SetShipLocations(pos1)
			b.SetShipLocations(pos2)
			return true
		}
	} else {
		return false
	}

}

func SelectTargets(play *types.Player, attackedBoard *types.GameBoard, aiBoard *types.GameBoard) {
	var x int
	var y int
	fmt.Print("Select x value")
	fmt.Scanln(&x)
	fmt.Print("Select y value")
	fmt.Scanln(&y)
	pos := types.NewPosition(x, y)
	hit := play.MakeMove(pos, aiBoard)
	attackedBoard.SetValue(pos, hit)
	if hit == true {
		fmt.Print("Hit")
	} else {
		fmt.Print("Miss")
	}
}

func CheckWin(p *types.Player) bool {
	if p.GetRemainingPieces() == 0 {
		return true
	} else {
		return false
	}
}

func RunGame(player *types.Player, ai *types.Player, pBoard *types.GameBoard, aiBoard *types.GameBoard, attackedBoard *types.GameBoard, state *types.Gamestate) {
	PickShipLocations(player, pBoard)
	for {
		pBoard.Displayboard()
		if state.GetTurn()%2 != 0 {
			SelectTargets(player, attackedBoard, aiBoard)
			CheckWin(ai)
		} else {
			//AIattack()
			CheckWin(player)
		}

		break

	}
}

func main() {
	state := types.NewGameState()
	player := types.NewPlayer("Zach", 12)
	aiPlayer := types.NewPlayer("AI", 12)
	playerBoard := types.NewGameBoard(length, width, player)
	attackedBoard := types.NewGameBoard(length, width, aiPlayer)
	aiBoard := types.NewGameBoard(width, length, aiPlayer)

	RunGame(player, aiPlayer, playerBoard, aiBoard, attackedBoard, state)
	fmt.Print(player.GetRemainingPieces())

}
