package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	world := InputToWorld()

	var count int
	for Pour(world) {
		count++
	}
	fmt.Println(count)
}

func Pour(world World) bool {
	p := aoc.Point2D{X: 500, Y: 0}

	for p.Y < world.Height-1 {
		if q := p.Down(); world.InBoundsPoint(q) && world.GetPoint(q) == Empty {
			p = q
			continue
		}
		if q := p.Down().Left(); world.InBoundsPoint(q) && world.GetPoint(q) == Empty {
			p = q
			continue
		}
		if q := p.Down().Right(); world.InBoundsPoint(q) && world.GetPoint(q) == Empty {
			p = q
			continue
		}

		world.SetPoint(p, Sand)
		break
	}

	return p.Y < world.Height-1
}

const (
	Empty = iota
	Wall
	Sand
)

type World struct {
	aoc.Grid2D[int]
}

func InputToWorld() World {
	var walls aoc.Set[aoc.Point2D]
	for _, line := range aoc.InputToLines(2022, 14) {
		points := strings.Split(line, " -> ")

		current := ParsePoint(points[0])
		walls.Add(current)

		for _, s := range points[1:] {
			end := ParsePoint(s)

			dx, dy := aoc.Sign(end.X-current.X), aoc.Sign(end.Y-current.Y)
			for current != end {
				current.X += dx
				current.Y += dy
				walls.Add(current)
			}
		}
	}

	_, br := aoc.GetBounds(walls.Entries())

	grid := aoc.NewGrid2D[int](br.X+1, br.Y+1)
	for p := range walls {
		grid.SetPoint(p, Wall)
	}

	return World{grid}
}

func ParsePoint(s string) aoc.Point2D {
	var p aoc.Point2D
	_, _ = fmt.Sscanf(s, "%d,%d", &p.X, &p.Y)
	return p
}
