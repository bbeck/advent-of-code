package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	points, velocities := InputToPoints(), InputToVelocities()

	var tl, br aoc.Point2D
	var tm int
	for tm = 1; ; tm++ {
		for i := 0; i < len(points); i++ {
			points[i] = aoc.Point2D{
				X: points[i].X + velocities[i].X,
				Y: points[i].Y + velocities[i].Y,
			}
		}

		tl, br = aoc.GetBounds(points)
		if br.Y-tl.Y <= 10 { // Characters are 10 pixels high
			break
		}
	}

	fmt.Println(tm)
}

func InputToPoints() []aoc.Point2D {
	var unused int
	return aoc.InputLinesTo(2018, 10, func(line string) (aoc.Point2D, error) {
		var p aoc.Point2D
		_, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &p.X, &p.Y, &unused, &unused)
		return p, err
	})
}

func InputToVelocities() []aoc.Point2D {
	var unused int
	return aoc.InputLinesTo(2018, 10, func(line string) (aoc.Point2D, error) {
		var p aoc.Point2D
		_, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &unused, &unused, &p.X, &p.Y)
		return p, err
	})
}
