package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		fields := strings.Fields(scanner.Text())
		K, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("atoi K %s", err)
		}
		C, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("atoi S %s", err)
		}
		S, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalf("atoi C %s", err)
		}

		fmt.Printf("Case #%d: %s\n", t+1, d(K, C, S))
	}
}

func d(K, C, S int) string {
	// return up to S positions
	// each position can check up to C tiles of the original sequence
	// we need to check all K tiles

	if S*C < K {
		return "IMPOSSIBLE"
	}

	numbers := []string{}

	n, b := 0, 0
	for k := 0; k < K; k++ {
		n = n*K + k
		b++
		if b >= C {
			numbers = append(numbers, strconv.Itoa(n+1))
			n, b = 0, 0
		}
	}
	if b > 0 {
		numbers = append(numbers, strconv.Itoa(n+1))
	}

	return strings.Join(numbers, " ")
}
