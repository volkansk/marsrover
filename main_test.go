package main

import (
	"testing"

	"github.com/volkansk/marsrover/utilities"
)

func TestModFunc(t *testing.T) {
	if utilities.Mod(4, 4) != 0 {
		t.Fatal("Test fail!")
	}
}

func TestSuccess(t *testing.T) {
	plateau, rover1, firstCommands, rover2, secondCommands := utilities.TryParseAllInputs("5 5", "1 2 N ", "LMLMLMLMM", "3 3 E ", "MMRMMRMRRM")
	if (plateau.XMax != 5 || plateau.YMax != 5) && rover1.X != 1 && rover2.Y != 3 && len(firstCommands) != 9 && len(secondCommands) != 10 {
		t.Errorf("expected to be successful")
	}
}

// tests will be written

// i hope so :)
