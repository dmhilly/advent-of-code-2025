package main

import (
	"bufio"
	"fmt"
	"os"
)

func largestTwoDigitSubseq(s string) int {
	maxFirst := s[0]
	idx := 0
	for i := 1; i < len(s)-1; i++ {
		if s[i] > maxFirst {
			maxFirst = s[i]
			idx = i
		}
	}
	maxSecond := s[idx+1]
	for i := idx + 2; i < len(s); i++ {
		if s[i] > maxSecond {
			maxSecond = s[i]
		}
	}
	return int(maxFirst-'0')*10 + int(maxSecond-'0')
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("open input.txt: %v\n", err)
		return
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		sum += largestTwoDigitSubseq(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("read input.txt: %v\n", err)
		return
	}
	fmt.Println("total:", sum)
}
