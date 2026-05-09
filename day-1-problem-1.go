package main

import (
	"fmt"
	"strconv"
)

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func main() {
	instructions := []string{
		"L68", "L30", "R48", "L5", "R60",
		"L55", "L1", "L99", "R14", "L82",
	}

	pos := 50
	zeros := 0

	for _, ins := range instructions {
		dir := ins[0]
		n, err := strconv.Atoi(ins[1:])
		if err != nil {
			fmt.Printf("bad instruction %q: %v\n", ins, err)
			return
		}

		if dir == 'L' {
			pos = mod(pos-n, 100)
		} else {
			pos = mod(pos+n, 100)
		}

		fmt.Println(pos)
		if pos == 0 {
			zeros++
		}
	}

	fmt.Printf("zeros seen: %d\n", zeros)
}
