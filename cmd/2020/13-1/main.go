package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
	"strings"
)

func main() {
	tm, buses := InputToBuses()

	bestWait := math.MaxInt
	bestBus := 0
	for _, bus := range buses {
		wait := (bus - tm%bus) % bus
		if wait < bestWait {
			bestWait = wait
			bestBus = bus
		}
	}

	fmt.Println(bestWait * bestBus)
}

func InputToBuses() (int, []int) {
	lines := aoc.InputToLines(2020, 13)
	tm := aoc.ParseInt(lines[0])

	var buses []int
	for _, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			buses = append(buses, aoc.ParseInt(s))
		}
	}

	return tm, buses
}
