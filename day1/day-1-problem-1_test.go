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
	wantZeros := 6

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

func TestRunInstructionsEdgeCases(t *testing.T) {
	// runInstructions always starts at pos=50. Cases that need a different
	// starting position prefix a setup move; setup crossings are folded into
	// wantZeros.
	tests := []struct {
		name          string
		instructions  []string
		wantPositions []int
		wantZeros     int
	}{
		{
			name:          "R lands exactly on 0",
			instructions:  []string{"R50"},
			wantPositions: []int{0},
			wantZeros:     1,
		},
		{
			name:          "L lands exactly on 0",
			instructions:  []string{"L50"},
			wantPositions: []int{0},
			wantZeros:     1,
		},
		{
			name:          "R wraps twice and lands on 0",
			instructions:  []string{"R150"},
			wantPositions: []int{0},
			wantZeros:     2,
		},
		{
			name:          "L wraps twice and lands on 0",
			instructions:  []string{"L150"},
			wantPositions: []int{0},
			wantZeros:     2,
		},
		{
			name:          "from 0, L1 walks away (no crossing)",
			instructions:  []string{"R50", "L1"},
			wantPositions: []int{0, 99},
			wantZeros:     1,
		},
		{
			name:          "from 0, R1 walks away (no crossing)",
			instructions:  []string{"R50", "R1"},
			wantPositions: []int{0, 1},
			wantZeros:     1,
		},
		{
			name:          "from 0, L100 laps back to 0",
			instructions:  []string{"R50", "L100"},
			wantPositions: []int{0, 0},
			wantZeros:     2,
		},
		{
			name:          "from 0, L101 crosses -100 only",
			instructions:  []string{"R50", "L101"},
			wantPositions: []int{0, 99},
			wantZeros:     2,
		},
		{
			name:          "from 0, R100 laps back to 0",
			instructions:  []string{"R50", "R100"},
			wantPositions: []int{0, 0},
			wantZeros:     2,
		},
		{
			name:          "R1000 from 50 crosses 10 hundreds",
			instructions:  []string{"R1000"},
			wantPositions: []int{50},
			wantZeros:     10,
		},
		{
			name:          "L1000 from 50 crosses 10 hundreds",
			instructions:  []string{"L1000"},
			wantPositions: []int{50},
			wantZeros:     10,
		},
		{
			name:          "R1000 from 0 crosses 10 hundreds",
			instructions:  []string{"R50", "R1000"},
			wantPositions: []int{0, 0},
			wantZeros:     11,
		},
		{
			name:          "L1000 from 0 crosses 10 hundreds",
			instructions:  []string{"R50", "L1000"},
			wantPositions: []int{0, 0},
			wantZeros:     11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			positions, zeros, err := runInstructions(tt.instructions)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(positions, tt.wantPositions) {
				t.Errorf("positions: got %v, want %v", positions, tt.wantPositions)
			}
			if zeros != tt.wantZeros {
				t.Errorf("zeros: got %d, want %d", zeros, tt.wantZeros)
			}
		})
	}
}
