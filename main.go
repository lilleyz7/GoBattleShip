package main

import "fmt"

const length = 8
const width = 8

func PickShipLocations(p *Player, b *GameBoard) {
	var xVal int
	var yVal int
	var alignment string
	i := 0

	for i < 5 {
		fmt.Println("You will select 4 points on the board and select to face the ships vert or horizontally")
		fmt.Println("Select your first x value:")
		fmt.Scanln(&xVal)
		fmt.Println("Select your first y value:")
		fmt.Scanln(&yVal)
		fmt.Println("Vertical or horizontal? Enter v or h")
		fmt.Scanln(&alignment)
		SetLocation(xVal, yVal, alignment, b)
	}

}

func SetLocation(x int, y int, a string, b *GameBoard) {
	b.board[x][y] = "S"

	if a == "v" {
		b.board[x][y+1] = "S"
		b.board[x][y-1] = "S"
	} else if a == "h" {
		b.board[x+1][y] = "S"
		b.board[x-1][y] = "S"
	} else {
		fmt.Println("Invalid entry")
	}
}

func main() {
	//state := NewGameState()
	player := NewPlayer("Zach")
	//aiPlayer := NewPlayer("AI")
	playerBoard := NewGameBoard(length, width)
	//aiBoard := NewGameBoard(width, length)
	PickShipLocations(player, playerBoard)

}
