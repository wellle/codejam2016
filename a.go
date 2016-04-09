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
		N, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("atoi N %s", err)
		}

		fmt.Printf("Case #%d: %s\n", t+1, a(N))
	}
}

func a(N int) string {
	if N == 0 {
		return "INSOMNIA"
	}

	chars := map[int]struct{}{}

	sum := N
	for {
		setChars(chars, sum)

		if len(chars) >= 10 {
			return strconv.Itoa(sum)
		}

		sum += N
	}
}

func setChars(chars map[int]struct{}, n int) {
	for n > 0 {
		c := n % 10
		chars[c] = struct{}{}
		n /= 10
	}
}
