package main

import (
	"github.com/volkansk/marsrover/actions"
	"github.com/volkansk/marsrover/utilities"
)

func main() {
	plateau, rover1, firstCommands, rover2, secondCommands := utilities.ReadAndValidateLines()

	utilities.R1 = actions.SendCommandsToRover(*plateau, *rover1, firstCommands, 1)
	utilities.R2 = actions.SendCommandsToRover(*plateau, *rover2, secondCommands, 2)

	utilities.ReturnRoverCoordinates(utilities.R1)
	utilities.ReturnRoverCoordinates(utilities.R2)

	utilities.PressAnyKeyToExit()
}
