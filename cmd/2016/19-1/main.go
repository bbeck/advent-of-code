package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	n := aoc.InputToInt(2016, 19)

	var elves aoc.Ring[*Elf]
	for i := 0; i < n; i++ {
		elves.InsertAfter(&Elf{ID: i + 1, Presents: 1})
	}
	elves.Next()

	for i := 0; i < n-1; i++ {
		current := elves.Current()
		next := elves.Next()
		current.Presents += next.Presents
		elves.Remove()
	}

	fmt.Println(elves.Current().ID)
}

type Elf struct {
	ID       int
	Presents int
}
