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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("open input.txt: %v\n", err)
		return
	}
	defer f.Close()

	pos := 50
	zeros := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ins := scanner.Text()
		if ins == "" {
			continue
		}

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
	if err := scanner.Err(); err != nil {
		fmt.Printf("read input.txt: %v\n", err)
		return
	}

	fmt.Printf("zeros seen: %d\n", zeros)
}
