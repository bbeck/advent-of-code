package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var digits []int
	for _, c := range aoc.InputToString(2015, 10) {
		digits = append(digits, aoc.ParseInt(string(c)))
	}

	for i := 0; i < 40; i++ {
		digits = LookAndSay(digits)
	}

	fmt.Println(len(digits))
}

func LookAndSay(s []int) []int {
	var output []int

	last, count := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == last {
			count++
			continue
		}

		output = append(output, []int{count, last}...)
		last = s[i]
		count = 1
	}

	output = append(output, []int{count, last}...)
	return output
}
