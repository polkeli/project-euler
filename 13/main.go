package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	w int = 50
	d int = 100
)

func index(x, y int) int {
	return (y*w)+x
}

func fill(n string) []int {
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		m = make([]int, w*d)
		offset = 0
		scanner = bufio.NewScanner(f)
	)

	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		for i, v := range line {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			m[index(i, offset)] = n
		}
		offset++
	}

	return m
}

func main() {
	m := fill("./digits.txt")

	var (
		total = ""
		carried = 0
	)

	for x := w-1; x >= 0; x-- {
		res := carried
		for y := 0; y < d; y++ {
			res += m[index(x, y)]
		}

		if n := fmt.Sprint(res); len(n) > 1 {
			var err error
			carried, err = strconv.Atoi(n[0:len(n)-1])
			if err != nil {
				panic(err)
			}

			total = fmt.Sprintf("%v%s", n[len(n)-1:], total)
		}else {
			total = fmt.Sprintf("%v%s", res, total)
		}
	}

	fmt.Printf("full:		%s\n", total)
	fmt.Printf("first 10:	%s\n", string(total[0:10]))
}
