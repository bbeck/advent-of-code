package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	caves := InputToCaves()
	count := CountPaths("start", "end", caves, aoc.SetFrom("start"), false)
	fmt.Println(count)
}

func CountPaths(current, goal string, caves map[string][]string, seen aoc.Set[string], used bool) int {
	if current == goal {
		return 1
	}

	var count int
	for _, n := range caves[current] {
		isLower := n == strings.ToLower(n)
		if n == "start" || (isLower && seen.Contains(n) && used) {
			continue
		}

		count += CountPaths(n, goal, caves, seen.UnionElems(n), used || (isLower && seen.Contains(n)))
	}

	return count
}

func InputToCaves() map[string][]string {
	caves := make(map[string][]string)
	for _, line := range aoc.InputToLines(2021, 12) {
		lhs, rhs, _ := strings.Cut(line, "-")
		caves[lhs] = append(caves[lhs], rhs)
		caves[rhs] = append(caves[rhs], lhs)
	}
	return caves
}
