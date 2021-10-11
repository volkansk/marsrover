package actions

import (
	"github.com/volkansk/marsrover/enums"
	"github.com/volkansk/marsrover/exceptions"
	"github.com/volkansk/marsrover/models"
	"github.com/volkansk/marsrover/utilities"
)

func SendCommandsToRover(p models.Plateau, r models.Rover, c []string, o int) models.Rover {
	for _, element := range c {
		switch element {
		case "M":
			r = MoveRover(r, p)
			//coordinate of other rover is checked
			if o == 1 && r.X == utilities.R2.X && r.Y == utilities.R2.Y {
				exceptions.ThrowRoverCollide()
			} else if o == 2 && r.X == utilities.R1.X && r.Y == utilities.R1.Y {
				exceptions.ThrowRoverCollide()
			}
			continue
		case "R":
			r.Direction = utilities.Mod(r.Direction+1, 4)
			continue
		case "L":
			r.Direction = utilities.Mod(r.Direction-1, 4)
			continue
		default:
			exceptions.ThrowInvalidCommand()
		}
	}
	return r
}

func MoveRover(r models.Rover, p models.Plateau) models.Rover {
	switch r.Direction {
	case int(enums.N):
		if r.Y == p.YMax {
			exceptions.ThrowPlateauAreaExceeded()
		}
		r.Y = r.Y + 1
		return r
	case int(enums.E):
		if r.X == p.XMax {
			exceptions.ThrowPlateauAreaExceeded()
		}
		r.X = r.X + 1
		return r
	case int(enums.S):
		if r.Y == 0 {
			exceptions.ThrowPlateauAreaExceeded()
		}
		r.Y = r.Y - 1
		return r
	case int(enums.W):
		if r.X == 0 {
			exceptions.ThrowPlateauAreaExceeded()
		}
		r.X = r.X - 1
		return r
	}
	return r
}
