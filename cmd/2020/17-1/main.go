package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	cube := InputToCube()
	for i := 0; i < 6; i++ {
		cube = Next(cube)
	}

	fmt.Println(len(cube))
}

func Next(cube aoc.Set[aoc.Point3D]) aoc.Set[aoc.Point3D] {
	var next aoc.Set[aoc.Point3D]

	tl, br := aoc.GetBounds3D(cube.Entries())
	for x := tl.X - 1; x <= br.X+1; x++ {
		for y := tl.Y - 1; y <= br.Y+1; y++ {
			for z := tl.Z - 1; z <= br.Z+1; z++ {
				p := aoc.Point3D{X: x, Y: y, Z: z}

				var active int
				for _, n := range p.Neighbors() {
					if cube.Contains(n) {
						active++
					}
				}

				if cube.Contains(p) && (active == 2 || active == 3) {
					next.Add(p)
				} else if !cube.Contains(p) && (active == 3) {
					next.Add(p)
				}
			}
		}
	}

	return next
}

func InputToCube() aoc.Set[aoc.Point3D] {
	var cube aoc.Set[aoc.Point3D]
	for y, line := range aoc.InputToLines(2020, 17) {
		for x, c := range line {
			if c == '#' {
				cube.Add(aoc.Point3D{X: x, Y: y})
			}
		}
	}

	return cube
}
