package types

type Ship struct {
	location Position
	sunk     bool
}

// needs hella work
func (s *Ship) DestroyShip() {
	s = &Ship{}

}

func NewShip(pos Position) *Ship {
	newShip := &Ship{
		location: pos,
		sunk:     false,
	}
	return newShip
}
