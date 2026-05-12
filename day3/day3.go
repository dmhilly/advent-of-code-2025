package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func largestKDigitSubseq(s string, k int) int {
	l := len(s)
	skipsRemaining := l - k
	idx := 0
	var ret string

	for len(ret) < k {
		var largestNum byte = '0'
		largestIdx := idx
		for i := idx; i <= idx+skipsRemaining; i++ {
			if s[i] > largestNum {
				largestNum = s[i]
				largestIdx = i
			}
		}

		ret += string(largestNum)
		idx = largestIdx + 1
		skipsRemaining = l - k - (idx - len(ret))
	}

	n, err := strconv.Atoi(ret)
	if err != nil {
		panic(err)
	}
	return n
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
		sum += largestKDigitSubseq(line, 12)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("read input.txt: %v\n", err)
		return
	}
	fmt.Println("total:", sum)
}
