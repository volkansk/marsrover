package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/volkansk/marsrover/actions"
	"github.com/volkansk/marsrover/utilities"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plateau area (like -> 5 5): ")
	input1, _ := reader.ReadString('\n')

	fmt.Print("Enter first rover coordinates (like -> 1 2 N): ")
	input2, _ := reader.ReadString('\n')

	fmt.Print("Enter first rover commands (like -> LMLMLMLMM): ")
	input3, _ := reader.ReadString('\n')

	fmt.Print("Enter second rover coordinates(like -> 3 3 E): ")
	input4, _ := reader.ReadString('\n')

	fmt.Print("Enter second rover commands(like -> MMRMMRMRRM): ")
	input5, _ := reader.ReadString('\n')

	plateau, rover1, firstCommands, rover2, secondCommands := utilities.TryParseAllInputs(input1, input2, input3, input4, input5)

	utilities.P = *plateau
	utilities.R1 = *rover1
	utilities.R2 = *rover2

	utilities.R1 = actions.SendCommandsToRover(*plateau, *rover1, firstCommands, 1)
	utilities.R2 = actions.SendCommandsToRover(*plateau, *rover2, secondCommands, 2)

	utilities.ReturnRoverCoordinates(utilities.R1)
	utilities.ReturnRoverCoordinates(utilities.R2)

	utilities.PressAnyKeyToExit()
}
