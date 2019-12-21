package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	reactions := make(map[string]Reaction)
	for _, reaction := range InputToReactions(2019, 14) {
		reactions[reaction.output] = reaction
	}

	// Perform a binary search to figure out the maximum amount of fuel that we
	// can generate from our ore.
	L := 0
	R := 1_000_000_000
	for L <= R {
		target := (L + R) / 2
		needed := Run(reactions, target)

		if needed < 1_000_000_000_000 {
			L = target + 1
		} else {
			R = target - 1
		}
	}

	fmt.Printf("maximum fuel: %d\n", R)
}

func Run(reactions map[string]Reaction, needed int) int {
	// These are the chemicals we have and the quantities that we have them in,
	// we want to keep breaking these down until the only positive quantities that
	// we have is for ORE.
	chemicals := map[string]int{
		"FUEL": needed,
	}

	// Choose a chemical to break down, if there are none left then the empty
	// string is returned.
	choose := func(chemicals map[string]int) string {
		for chemical, quantity := range chemicals {
			if chemical != "ORE" && quantity > 0 {
				return chemical
			}
		}

		return ""
	}

	for {
		chemical := choose(chemicals)
		if chemical == "" {
			break
		}

		reaction := reactions[chemical]

		multiple := chemicals[chemical] / reaction.quantity
		if multiple == 0 {
			multiple = 1
		}

		chemicals[chemical] -= reaction.quantity * multiple
		for input, quantity := range reaction.inputs {
			chemicals[input] += quantity * multiple
		}
	}

	return chemicals["ORE"]
}

type Reaction struct {
	inputs   map[string]int
	output   string
	quantity int
}

func InputToReactions(year, day int) []Reaction {
	var reactions []Reaction
	for _, line := range aoc.InputToLines(year, day) {
		sides := strings.Split(line, " => ")

		inputs := make(map[string]int)
		for _, part := range strings.Split(sides[0], ", ") {
			var quantity int
			var chemical string
			if _, err := fmt.Sscanf(part, "%d %s", &quantity, &chemical); err != nil {
				log.Fatalf("unable to parse input part: %s", part)
			}

			inputs[chemical] = quantity
		}

		var quantity int
		var output string
		if _, err := fmt.Sscanf(sides[1], "%d %s", &quantity, &output); err != nil {
			log.Fatalf("unable to parse output part: %s", sides[1])
		}

		reactions = append(reactions, Reaction{inputs, output, quantity})
	}

	return reactions
}