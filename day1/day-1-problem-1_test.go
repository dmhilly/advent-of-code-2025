package main

import (
	"reflect"
	"testing"
)

func TestRunInstructionsExample(t *testing.T) {
	instructions := []string{
		"L68", "L30", "R48", "L5", "R60",
		"L55", "L1", "L99", "R14", "L82",
	}
	wantPositions := []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
	wantZeros := 3

	positions, zeros, err := runInstructions(instructions)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(positions, wantPositions) {
		t.Errorf("positions: got %v, want %v", positions, wantPositions)
	}
	if zeros != wantZeros {
		t.Errorf("zeros: got %d, want %d", zeros, wantZeros)
	}
}
