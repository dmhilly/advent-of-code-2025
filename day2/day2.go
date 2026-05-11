package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "18623-26004,226779-293422,65855-88510,868-1423,248115026-248337139,903911-926580,97-121,67636417-67796062,24-47,6968-10197,193-242,3769-5052,5140337-5233474,2894097247-2894150301,979582-1016336,502-646,9132195-9191022,266-378,58-91,736828-868857,622792-694076,6767592127-6767717303,2920-3656,8811329-8931031,107384-147042,941220-969217,3-17,360063-562672,7979763615-7979843972,1890-2660,23170346-23308802"
	// testInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	totalSum, err := calculateTotalSum(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("total: ", totalSum)
}

func calculateTotalSum(input string) (int, error) {
	var totalSum int
	ranges, err := parseRanges(input)
	if err != nil {
		return 0, err
	}

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if isInvalidID(i) {
				totalSum += i
			}
		}
	}

	return totalSum, nil
}

func isInvalidID(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	if l%2 == 1 {
		return false
	}
	return s[:l/2] == s[l/2:]
}

type IDRange struct {
	Start, End int
}

func parseRanges(input string) ([]IDRange, error) {
	parts := strings.Split(input, ",")
	ranges := make([]IDRange, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		dash := strings.Index(p, "-")
		if dash < 0 {
			return nil, fmt.Errorf("bad range %q: missing '-'", p)
		}
		start, err := strconv.Atoi(p[:dash])
		if err != nil {
			return nil, fmt.Errorf("bad range %q start: %w", p, err)
		}
		end, err := strconv.Atoi(p[dash+1:])
		if err != nil {
			return nil, fmt.Errorf("bad range %q end: %w", p, err)
		}
		ranges = append(ranges, IDRange{Start: start, End: end})
	}
	return ranges, nil
}
