package types

type Gamestate struct {
	turn       int
	activeTurn Player
	isRunning  bool
}

func NewGameState() *Gamestate {
	return &Gamestate{
		turn:      0,
		isRunning: true,
	}
}
