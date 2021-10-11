package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/volkansk/marsrover/enums"
	"github.com/volkansk/marsrover/exceptions"
	"github.com/volkansk/marsrover/models"
)

var R1 models.Rover
var R2 models.Rover
var P models.Plateau

func Mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func TryParseDirection(input string) (d enums.Direction) {
	switch input {
	case "N":
		return enums.N
	case "E":
		return enums.E
	case "S":
		return enums.S
	case "W":
		return enums.W
	default:
		exceptions.ThrowInvalidDirection()
		return
	}
}

func TryParseCommandType(input string) (c enums.CommandType) {
	switch input {
	case enums.M.String():
		return
	case enums.R.String():
		return
	case enums.L.String():
		return
	default:
		exceptions.ThrowInvalidCommand()
		return
	}
}

func TryParsePlateau(input string) (p *models.Plateau) {
	split := strings.Split(strings.TrimSpace(input), " ")
	if len(split) != 2 {
		exceptions.ThrowMoreOrLessInput()
	}
	px, err := strconv.Atoi(split[0])
	if err != nil {
		exceptions.ThrowMustNumeric()
	}
	py, err := strconv.Atoi(split[1])
	if err != nil {
		exceptions.ThrowMustNumeric()
	}
	if px < 1 || py < 1 {
		exceptions.ThrowMustNumeric()
	}
	return models.NewPlateau(px, py)
}

func TryParseRover(input string) (r *models.Rover) {
	split := strings.Split(strings.TrimSpace(input), " ")
	if len(split) != 3 {
		exceptions.ThrowMoreOrLessInput()
	}
	r1x, err := strconv.Atoi(split[0])
	if err != nil {
		exceptions.ThrowMustNumeric()
	}
	r1y, err := strconv.Atoi(split[1])
	if err != nil {
		exceptions.ThrowMustNumeric()
	}
	r1d := TryParseDirection(split[2])
	return models.NewRover(r1x, r1y, int(r1d))
}

func TryParseCommand(input string) []string {
	split := strings.Split(strings.TrimSpace(input), "")
	for _, x := range split {
		TryParseCommandType(x)
	}
	return split
}

func ReturnRoverCoordinates(r models.Rover) {
	fmt.Println(strconv.Itoa(r.X), strconv.Itoa(r.Y), enums.Direction(r.Direction).String())
}

func PressAnyKeyToExit() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Press Any Key To Exit...")
	input1, _ := reader.ReadString('\n')
	fmt.Println(input1)
}

func ReadAndValidateLines() (*models.Plateau, *models.Rover, []string, *models.Rover, []string) {

	reader := bufio.NewReader(os.Stdin)
	// plateau area
	fmt.Print("Enter plateau area (like -> 5 5): ")
	input1, _ := reader.ReadString('\n')
	P := TryParsePlateau(input1)

	// rover coordinates
	fmt.Print("Enter first rover coordinates (like -> 1 2 N): ")
	input2, _ := reader.ReadString('\n')
	R1 := TryParseRover(input2)
	if R1.X > P.XMax || R1.Y > P.YMax || R1.X < 0 || R1.Y < 0 {
		exceptions.ThrowOutOfArea()
	}

	// first rover commands
	fmt.Print("Enter first rover commands (like -> LMLMLMLMM): ")
	input3, _ := reader.ReadString('\n')
	c1 := TryParseCommand(input3)

	// second rover coordinates
	fmt.Print("Enter second rover coordinates(like -> 3 3 E): ")
	input4, _ := reader.ReadString('\n')
	R2 := TryParseRover(input4)
	if R2.X > P.XMax || R2.Y > P.YMax || R2.X < 0 || R2.Y < 0 {
		exceptions.ThrowOutOfArea()
	}

	// second rover commands
	fmt.Print("Enter second rover commands(like -> MMRMMRMRRM): ")
	input5, _ := reader.ReadString('\n')
	c2 := TryParseCommand(input5)

	return P, R1, c1, R2, c2
}
