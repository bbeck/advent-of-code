package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 5),
		Input:  func() int { return 5 },
		Output: func(value int) { fmt.Println(value) },
	}
	cpu.Execute()
}
