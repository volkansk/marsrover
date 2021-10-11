package exceptions

import (
	"errors"
	"log"
)

func ThrowRoverCollide() {
	log.Fatal(errors.New("movement stopped as rovers collide"))
}

func ThrowPlateauAreaExceeded() {
	log.Fatal(errors.New("specified plateau area exceeded"))
}

func ThrowInvalidCommand() {
	log.Fatal(errors.New("command is invalid. he possible letters are L, R and M"))
}

func ThrowMoreOrLessInput() {
	log.Fatal(errors.New("there are more or less inputs than expected"))
}

func ThrowMustNumeric() {
	log.Fatal(errors.New("numeric value must be entered"))
}

func ThrowOutOfArea() {
	log.Fatal(errors.New("coordinates must be entered within the plateau area"))
}

func ThrowInvalidDirection() {
	log.Fatal(errors.New("direction is invalid. he possible letters are N, E, S and W"))
}
