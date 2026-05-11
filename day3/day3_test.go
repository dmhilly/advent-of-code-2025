package main

import "testing"

func TestLargestTwoDigitSubseq(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	sum := 0
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := largestTwoDigitSubseq(tt.s)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
		sum += tt.want
	}

	if sum != 357 {
		t.Errorf("example sum: got %d, want 357", sum)
	}
}
