package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, group := range InputToGroups(2020, 6) {
		intersection := group[0]
		for _, answers := range group {
			intersection = intersection.Intersect(answers)
		}

		count += intersection.Size()
	}

	fmt.Println(count)
}

type Group []aoc.Set

func InputToGroups(year, day int) []Group {
	var groups []Group

	var current Group
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			groups = append(groups, current)
			current = make(Group, 0)
			continue
		}

		// Each line is a single person's answers
		answers := aoc.NewSet()
		for _, question := range line {
			answers.Add(string(question))
		}
		current = append(current, answers)
	}
	groups = append(groups, current)

	return groups
}
