package types

type Position struct {
	xVal int
	yVal int
}

func NewPosition(x, y int) *Position {
	return &Position{
		xVal: x,
		yVal: y,
	}
}

func (p *Position) UpdateVert() (*Position, *Position) {
	v1 := NewPosition(p.xVal, p.yVal+1)
	v2 := NewPosition(p.xVal, p.yVal-1)
	return v1, v2
}

func (p *Position) UpdateHoriz() (*Position, *Position) {
	v1 := NewPosition(p.xVal-1, p.yVal)
	v2 := NewPosition(p.xVal+1, p.yVal)
	return v1, v2
}

type Player struct {
	name            string
	selectedMoves   []*Position
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

func (pl *Player) AddMove(p *Position) {
	pl.selectedMoves = append(pl.selectedMoves, p)
}

func (pl *Player) MakeMove(p *Position, g *GameBoard) bool {
	for _, val := range pl.selectedMoves {
		if val == p {
			//already made move
			return false
		}
	}
	attackSuccess := g.TakeHit(p)
	pl.AddMove(p)
	return attackSuccess
}
