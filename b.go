package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("atoi T %s", err)
	}

	for t := 0; t < T; t++ {
		scanner.Scan()
		S := scanner.Text()

		fmt.Printf("Case #%d: %d\n", t+1, b(S))
	}
}

func b(S string) int {
	flips := 0
	last := ' '
	for _, c := range S {
		if last != ' ' && c != last {
			flips++
		}
		last = c
	}
	if last == '-' {
		flips++
	}
	return flips
}
