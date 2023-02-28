package main

type Ship struct {
	location []Position
	sunk     bool
}

type ShipStore struct {
	ships []Ship
}

// needs hella work
func (s *Ship) DestroyShip(store ShipStore) {
	s = &Ship{}

}

func NewShip(store ShipStore, pos []Position) *Ship {
	newShip := &Ship{
		location: pos,
		sunk:     false,
	}
	store.ships = append(store.ships, *newShip)
	return newShip
}
