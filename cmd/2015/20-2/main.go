package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	desired := aoc.InputToInt(2015, 20)

	houses := make([]int, desired+1)
	for elf := 1; elf <= desired; elf++ {
		for i := 1; i <= 50; i++ {
			house := elf * i
			if house <= desired {
				houses[house] += 11 * elf
			}
		}
	}

	var house int
	for house = 1; house <= desired; house++ {
		if houses[house] >= desired {
			break
		}
	}
	fmt.Println(house)
}
