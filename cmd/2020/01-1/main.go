package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 1)
	for i, a := range ns {
		for _, b := range ns[i+1:] {
			if a+b == 2020 {
				fmt.Println(a * b)
			}
		}
	}
}
