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
		N, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("atoi K %s", err)
		}
		J, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("atoi S %s", err)
		}

		fmt.Printf("Case #%d:\n", t+1)
		c(N, J)
	}
}

var (
	bases  = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	P      = 614889782588491410 // product of primes above
)

func c(N, J int) {
	iterate(N, J, []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, "", "1")
}

func iterate(left, limit int, values []int, s, next string) int {
	if limit == 0 {
		return 0
	}

	if left == 0 {
		factors := []int{}
		for i, _ := range bases {
			factors = append(factors, factor(values[i]))
			if factors[i] == 0 {
				return 0
			}
		}

		fmt.Printf("%s", s)
		for _, factor := range factors {
			fmt.Printf(" %d", factor)
		}
		fmt.Printf("\n")
		return 1
	}

	v := []int{}
	for i, b := range bases {
		v = append(v, values[i]*b%P)
		if next == "1" {
			v[i]++
		}
	}

	found := iterate(left-1, limit, v, s+next, "1")
	if left > 2 {
		found += iterate(left-1, limit-found, v, s+next, "0")
	}
	return found
}

func factor(n int) int {
	for _, p := range primes {
		if n%p == 0 {
			return p
		}
	}

	return 0
}
