package main

type Gamestate struct {
	turn       int
	activeTurn Player
	isRunning  bool
}

func NewGame() *Gamestate {
	player1 := NewPlayer("Zach")
	return &Gamestate{
		turn:       0,
		activeTurn: *player1,
		isRunning:  true,
	}
}
