package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	area := aoc.Make2D[int](1001, 1001)
	for _, claim := range InputToClaims() {
		for dx := 0; dx < claim.Width; dx++ {
			for dy := 0; dy < claim.Height; dy++ {
				area[claim.TL.Y+dy][claim.TL.X+dx]++
			}
		}
	}

	var count int
	for y := 0; y <= 1000; y++ {
		for x := 0; x <= 1000; x++ {
			if area[y][x] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Claim struct {
	ID            string
	TL            aoc.Point2D
	Width, Height int
}

func InputToClaims() []Claim {
	return aoc.InputLinesTo(2018, 3, func(line string) (Claim, error) {
		var claim Claim
		fmt.Sscanf(line, "#%s @ %d,%d: %dx%d", &claim.ID, &claim.TL.X, &claim.TL.Y, &claim.Width, &claim.Height)

		return claim, nil
	})
}
