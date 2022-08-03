package main

import (
	"fmt"
)

/*
* 2520 is the smallest number that can be divided by each of the numbers from
* 1 to 10 without any remainder.
*
* What is the smallest positive number that is evenly divisible by all of the
* numbers from 1 to 20?
*/

func main() {
	done := false
	for x := 2520; !done; x++{
		found := true
		for y := 1; y <= 20; y++ {
			if (x/y) % 2 != 0 {
				found = false
				break
			}
		}

		if found {
			fmt.Println(x)

			// bit of DEBUG output to confirm
			/*
			* for y := 1; y <= 20; y++ {
			*	fmt.Printf("(%d/%d) %% 2 = %d\n",x, y, (x/y) % 2)
			* }
			*/
			done = found // this breaks us out of the loop
		}
	}
}
