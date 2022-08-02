package main

import (
	"fmt"
	"strings"
	"time"
)

func even(n int) int {
	return n/2
}

func odd(n int) int {
	return 3*n+1
}

func chain(seed int) []int {
	vals := []int{seed}
	for {
		if n := vals[len(vals)-1]; n%2 == 0 {
			vals = append(vals, even(n))
		} else {
			vals = append(vals, odd(n))
		}

		if n := vals[len(vals)-1]; n < 1 {
			panic("negative numbers are band")
		} else if n == 1 {
			return vals
		}
	}
}

func pretty(n []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(n)), "->"), "[]")
}

func main() {
	const max = 1000000

	var (
		results = make(chan []int, 1)
		start = time.Now()
	)

	for i := 2; i < max; i++ {
		go func(val int) {
			results<-chain(val)
		}(i)
	}

	longest := []int{}
	for i := 2; i < max; i++{
		if r := <-results; len(longest) < len(r) {
			longest = make([]int, len(r))
			copy(longest, r)
		}
	}

	end := time.Since(start)
	//fmt.Println(pretty(longest))
	fmt.Printf("longest: %d\n", len(longest))
	fmt.Printf("starter: %d\n", longest[0])
	fmt.Printf("%vs\n", end.Seconds())
}
