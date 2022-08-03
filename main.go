package main

import (
	"fmt"
	"math"
)

/*
* 2520 is the smallest number that can be divided by each of the numbers from
* 1 to 10 without any remainder.
*
* What is the smallest positive number that is evenly divisible by all of the
* numbers from 1 to 20?
*/

func check(n float64) bool {
	for i := 1.0; i <= 20; i++ {
		//fmt.Printf("(%f/%f) %% 1 = %f\n", n, i, math.Mod(n/i, 1))
		if math.Mod(n/i, 1) != 0 {
			return false
		}
	}
	return true
}

func main() {
	for x := 1.0; true; x++{
		//fmt.Printf("checking %f\n", x)
		if check(x) {
			fmt.Println(x)

			/*
			// bit of DEBUG output to confirm
			for y := 1.0; y <= 10; y++ {
				fmt.Printf("(%f/%f) %% 1 = %f\n",x, y, math.Mod(x/y, 1))
			}
			*/

			return
		}
	}
}
