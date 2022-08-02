package main

import (
	"fmt"
)

/*
* Each new term in the Fibonacci sequence is generated by adding the previous
* two terms. By starting with 1 and 2, the first 10 terms will be:
*			1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
* By considering the terms in the Fibonacci sequence whose values do not exceed
* four million, find the sum of the even-valued terms.
*/

const max = 4000000

func fib(a, b int, evens chan int, done chan struct{}) {
	c := a + b

	// do not exceed 4000000
	if c > max {
		done <- struct{}{}
		return
	}

	// if it's even, send over for processing
	if c % 2 == 0 {
		evens <- c
	}

	// keep on fib'ing
	fib(b, c, evens, done)
}

func main() {
	var (
		evens = make(chan int)
		done = make(chan struct{})

		sum = 0
	)

	go fib(1,2, evens, done)

	for {
		select {
		case e := <-evens:
			sum += e
		case <-done:
			close(evens)
			close(done)
			fmt.Println(sum)
			return
		}
	}
}
