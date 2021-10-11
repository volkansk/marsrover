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

// tests will be written

// i hope so :)
