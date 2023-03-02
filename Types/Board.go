package types

import "fmt"

type GameBoard struct {
	length int
	width  int
	board  [][]string
}

func NewGameBoard(len, wid int) *GameBoard {
	board := make([][]string, len)
	for i := 0; i < len; i++ {
		for j := 0; j < wid; j++ {
			board[i] = append(board[i], "O")
		}
	}
	return &GameBoard{
		length: len,
		width:  wid,
		board:  board,
	}
}

func (b *GameBoard) GetBoard() [][]string {
	return b.board
}

func (b *GameBoard) SetLocation(xVal, yVal int) bool {
	if xVal > 0 && yVal > 0 {
		if xVal <= 7 && yVal <= 7 {
			b.board[xVal][yVal] = "S"
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

func (b *GameBoard) TakeHit(xVal, yVal int) bool {
	location := b.board[xVal][yVal]
	if location == "S" {
		location = "H"
		return true
	} else {
		location = "X"
		return false
	}
}

func (b *GameBoard) Displayboard() {
	fmt.Printf("	0	1	2	3	4")
	fmt.Printf("\n")
	fmt.Printf("\n")
	for i := 0; i < b.length; i++ {
		fmt.Printf("%d	", i)
		for j := 0; j < b.length; j++ {
			fmt.Printf("%s	", b.board[i][j])
		}
		fmt.Printf("\n")
	}
}
