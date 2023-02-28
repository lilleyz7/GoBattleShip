package main

type Position struct {
	xVal int
	yVal int
}

type Player struct {
	name          string
	currentBoard  GameBoard
	selectedMoves []Position
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
	}
}

func (pl *Player) AddMove(p Position) {
	pl.selectedMoves = append(pl.selectedMoves, p)
}

func (pl *Player) MakeMove(p Position, g GameBoard) bool {
	for _, val := range pl.selectedMoves {
		if val == p {
			//already made move
			return false
		}
	}
	pos := g.board[p.xVal][p.yVal]
	switch pos {
	case 0:
		pos = "X"
		pl.AddMove(p)
		return false
	case 1:
		pos = "H"
		pl.AddMove(p)
		return true
	default:
		//error
		return false
	}
}
