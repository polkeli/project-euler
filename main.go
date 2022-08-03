package main

import (
	"fmt"
)

/*
* A palindromic number reads the same both ways. The largest palindrome made
* from the product of two 2-digit numbers is 9009 = 91 × 99.
*
* Find the largest palindrome made from the product of two 3-digit numbers.
 */

func reverse(s string) string {
	r := []rune(s)
	for x, y := 0, len(r)-1; x < y; x, y = x+1, y-1 {
		r[x], r[y] = r[y], r[x]
	}

	return string(r)
}

func palindromic(n int) bool {
	s := fmt.Sprintf("%d", n)

	mid := len(s) / 2
	prefix := s[:mid]

	var suffix string
	if len(s) % 2 == 0 {
		// even length
		suffix = reverse(s[mid:])
	} else {
		// odd length
		suffix = reverse(s[mid+1:])
	}

	return prefix == suffix
}

func sample(n int, pals chan int, done chan struct{}) {
	for i := n; i < 1000; i++ {
		p := n * i
		if palindromic(p) {
			pals<-p
		}
	}
	done<-struct{}{}
}

func collect(pals chan int, done chan struct{}) int {
	max := 0
	for {
		select {
		case p := <-pals:
			// fmt.Println(p) // little bit of DEBUG output to confirmation
			if p > max {
				max = p
			}

		case <-done:
			return max
		}
	}
}

func main() {
	max := 0
	for i := 100; i < 1000; i++ {
		var (
			pals = make(chan int)
			done = make(chan struct{}, 1)
		)
		go sample(i, pals, done)
		if p := collect(pals, done); p > max {
			max = p
		}
	}

	fmt.Printf("final: %d\n", max)
}
