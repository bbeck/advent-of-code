package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 9)

	var target int
	for i := 25; i < len(ns); i++ {
		if !SumExists(ns[i-25:i], ns[i]) {
			target = ns[i]
			break
		}
	}

	start, end := FindRange(ns, target)
	min := aoc.Min(ns[start:end]...)
	max := aoc.Max(ns[start:end]...)
	fmt.Println(min + max)
}

func SumExists(ns []int, target int) bool {
	var seen aoc.Set[int]
	for _, n := range ns {
		if seen.Contains(target - n) {
			return true
		}
		seen.Add(n)
	}
	return false
}

func FindRange(ns []int, target int) (int, int) {
	cumsum := make([]int, len(ns))
	for i, sum := 0, 0; i < len(ns); i++ {
		sum += ns[i]
		cumsum[i] = sum
	}

	for start := 0; start < len(cumsum)-1; start++ {
		for end := start + 1; end < len(cumsum); end++ {
			if cumsum[end]-cumsum[start] == target {
				return start, end
			}
		}
	}
	return 0, 0
}
