package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func runInstructions(instructions []string) (positions []int, zeros int, err error) {
	pos := 50
	for _, ins := range instructions {
		if ins == "" {
			continue
		}
		dir := ins[0]
		n, atoiErr := strconv.Atoi(ins[1:])
		if atoiErr != nil {
			return nil, 0, fmt.Errorf("bad instruction %q: %w", ins, atoiErr)
		}

		if dir == 'L' {
			if n >= pos {
				zeros += (n - pos) / 100
				if pos > 0 {
					zeros++
				}
			}
			pos = mod(pos-n, 100)
		} else {
			zeros += (pos + n) / 100
			pos = mod(pos+n, 100)
		}
		positions = append(positions, pos)
	}
	return positions, zeros, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("open input.txt: %v\n", err)
		return
	}
	defer f.Close()

	var instructions []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("read input.txt: %v\n", err)
		return
	}

	positions, zeros, err := runInstructions(instructions)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range positions {
		fmt.Println(p)
	}
	fmt.Printf("zeros seen: %d\n", zeros)
}
