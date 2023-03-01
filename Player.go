package main

type Position struct {
	xVal int
	yVal int
}

type Player struct {
	name          string
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
	attackSuccess := g.TakeHit(p.xVal, p.yVal)
	return attackSuccess
}
