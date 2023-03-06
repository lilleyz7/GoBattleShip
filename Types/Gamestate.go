package types

type Gamestate struct {
	turn      int
	isRunning bool
}

func NewGameState() *Gamestate {
	return &Gamestate{
		turn:      0,
		isRunning: true,
	}
}

func (g *Gamestate) GetTurn() int {
	return g.turn
}
