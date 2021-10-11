package models

type Plateau struct {
	XMax int
	YMax int
}

func NewPlateau(x, y int) *Plateau {
	p := &Plateau{
		XMax: x,
		YMax: y,
	}
	return p
}
