package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2017, 4) {
		if IsValid(line) {
			count++
		}
	}
	fmt.Println(count)
}

func IsValid(s string) bool {
	var seen aoc.Set[string]
	for _, word := range strings.Fields(s) {
		if !seen.Add(word) {
			return false
		}
	}

	return true
}
