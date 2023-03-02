package types

type Position struct {
	xVal int
	yVal int
}

type Player struct {
	name            string
	selectedMoves   []Position
	remainingPieces int
}

func NewPlayer(name string, pieces int) *Player {
	return &Player{
		name:            name,
		remainingPieces: pieces,
	}
}

func (pl *Player) GetRemainingPieces() int {
	return pl.remainingPieces
}

func (pl *Player) RemovePiece() bool {
	pl.remainingPieces--
	if pl.remainingPieces < 1 {
		return false
	} else {
		return true
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
