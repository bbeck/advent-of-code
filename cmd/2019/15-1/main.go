package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	open, goal := Explore()

	children := func(p aoc.Point2D) []aoc.Point2D {
		var children []aoc.Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if open.Contains(child) {
				children = append(children, child)
			}
		}
		return children
	}

	isGoal := func(p aoc.Point2D) bool {
		return p == goal
	}

	if path, ok := aoc.BreadthFirstSearch(aoc.Origin2D, children, isGoal); ok {
		fmt.Println(len(path) - 1) // The path includes the starting point.
	}
}

var Headings = []aoc.Heading{aoc.Up, aoc.Down, aoc.Left, aoc.Right}
var Reverse = map[aoc.Heading]aoc.Heading{
	aoc.Up:    aoc.Down,
	aoc.Down:  aoc.Up,
	aoc.Left:  aoc.Right,
	aoc.Right: aoc.Left,
}

func Explore() (aoc.Set[aoc.Point2D], aoc.Point2D) {
	var open aoc.Set[aoc.Point2D]
	var goal aoc.Point2D

	robot := NewRobot()
	current := aoc.Origin2D

	var helper func()
	helper = func() {
		for _, heading := range Headings {
			status := robot.Move(heading)
			if status == 0 {
				continue
			}

			current = current.Move(heading)
			if status == 2 {
				goal = current
			}

			if open.Add(current) {
				helper()
			}

			robot.Move(Reverse[heading])
			current = current.Move(Reverse[heading])
		}
	}
	helper()

	return open, goal
}

type Robot struct {
	CPU      cpus.IntcodeCPU
	Commands chan int
	Status   chan int
}

func NewRobot() *Robot {
	commands := make(chan int)
	status := make(chan int)

	robot := &Robot{
		CPU: cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(2019, 15),
			Input:  func() int { return <-commands },
			Output: func(value int) { status <- value },
		},
		Commands: commands,
		Status:   status,
	}
	go robot.CPU.Execute()
	return robot
}

func (r *Robot) Move(h aoc.Heading) int {
	mapping := map[aoc.Heading]int{
		aoc.Up:    1,
		aoc.Down:  2,
		aoc.Left:  3,
		aoc.Right: 4,
	}

	r.Commands <- mapping[h]
	return <-r.Status
}
