package main

type GameBoard struct {
	length int
	width  int
	board  [][]any
}

func NewGameBoard(len, wid int) *GameBoard {
	var board [][]any
	for i := 0; i < len; i++ {
		for j := 0; j < wid; j++ {
			board[i] = append(board[i], 0)
			board[j] = append(board[j], 0)
		}
	}
	return &GameBoard{
		length: len,
		width:  wid,
		board:  board,
	}
}
