package models

type Rover struct {
	X         int
	Y         int
	Direction int
}

func NewRover(x, y, direction int) *Rover {
	r := &Rover{
		X:         x,
		Y:         y,
		Direction: direction,
	}
	return r
}
