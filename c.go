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
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	P      = 614889782588491410 // product of primes above
)

func c(N, J int) {
	iterate(N, J, 0, 0, 0, 0, 0, 0, 0, 0, 0, "", true)
}

func iterate(left, limit, v2, v3, v4, v5, v6, v7, v8, v9, v10 int, s string, one bool) int {
	if limit == 0 {
		return 0
	}

	if left == 0 {
		f2 := factor(v2)
		if f2 == 0 {
			return 0
		}

		f3 := factor(v3)
		if f3 == 0 {
			return 0
		}

		f4 := factor(v4)
		if f4 == 0 {
			return 0
		}

		f5 := factor(v5)
		if f5 == 0 {
			return 0
		}

		f6 := factor(v6)
		if f6 == 0 {
			return 0
		}

		f7 := factor(v7)
		if f7 == 0 {
			return 0
		}

		f8 := factor(v8)
		if f8 == 0 {
			return 0
		}

		f9 := factor(v9)
		if f9 == 0 {
			return 0
		}

		f10 := factor(v10)
		if f10 == 0 {
			return 0
		}

		fmt.Printf("%s %d %d %d %d %d %d %d %d %d\n", s, f2, f3, f4, f5, f6, f7, f8, f9, f10)
		return 1
	}

	v2 = v2 * 2 % P
	v3 = v3 * 3 % P
	v4 = v4 * 4 % P
	v5 = v5 * 5 % P
	v6 = v6 * 6 % P
	v7 = v7 * 7 % P
	v8 = v8 * 8 % P
	v9 = v9 * 9 % P
	v10 = v10 * 10 % P

	if one {
		v2++
		v3++
		v4++
		v5++
		v6++
		v7++
		v8++
		v9++
		v10++
		s += "1"
	} else {
		s += "0"
	}

	left--
	found := iterate(left, limit, v2, v3, v4, v5, v6, v7, v8, v9, v10, s, true)
	limit -= found
	if left > 1 {
		found += iterate(left, limit, v2, v3, v4, v5, v6, v7, v8, v9, v10, s, false)
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
