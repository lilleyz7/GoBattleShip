package types

import (
	"fmt"
)

type GameBoard struct {
	length int
	width  int
	board  [][]string
	player *Player
}

func NewGameBoard(len, wid int, p *Player) *GameBoard {
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
		player: p,
	}
}

func (b *GameBoard) GetBoard() [][]string {
	return b.board
}

func (b *GameBoard) SetShipLocations(p *Position) bool {
	if p.xVal > 0 && p.yVal > 0 {
		if p.xVal <= 5 && p.yVal <= 5 {
			b.board[p.xVal][p.yVal] = "S"
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

func (b *GameBoard) SetValue(p *Position, hit bool) {
	if hit == true {
		b.board[p.xVal][p.yVal] = "H"
	} else {
		b.board[p.xVal][p.yVal] = "X"
	}
}

func (b *GameBoard) TakeHit(p *Position) bool {
	location := b.board[p.xVal][p.yVal]
	if location == "S" {
		location = "H"
		b.player.RemovePiece()
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
