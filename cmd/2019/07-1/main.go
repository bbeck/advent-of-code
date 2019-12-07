package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	phases := []int{0, 1, 2, 3, 4}

	var best int
	aoc.EnumeratePermutations(5, func(perm []int) {
		ZtoA := make(chan int, 2)
		AtoB := make(chan int, 2)
		BtoC := make(chan int, 2)
		CtoD := make(chan int, 2)
		DtoE := make(chan int, 2)
		EtoT := make(chan int)

		// Initialize the phase settings
		ZtoA <- phases[perm[0]]
		AtoB <- phases[perm[1]]
		BtoC <- phases[perm[2]]
		CtoD <- phases[perm[3]]
		DtoE <- phases[perm[4]]

		ZtoA <- 0 // First amplifier's input is hardcoded to zero
		A := &CPU{
			memory: InputToMemory(2019, 7),
			input:  func(int) int { return <-ZtoA },
			output: func(value int) { AtoB <- value },
		}
		go A.Execute()

		B := &CPU{
			memory: InputToMemory(2019, 7),
			input:  func(int) int { return <-AtoB },
			output: func(value int) { BtoC <- value },
		}
		go B.Execute()

		C := &CPU{
			memory: InputToMemory(2019, 7),
			input:  func(int) int { return <-BtoC },
			output: func(value int) { CtoD <- value },
		}
		go C.Execute()

		D := &CPU{
			memory: InputToMemory(2019, 7),
			input:  func(int) int { return <-CtoD },
			output: func(value int) { DtoE <- value },
		}
		go D.Execute()

		E := &CPU{
			memory: InputToMemory(2019, 7),
			input:  func(int) int { return <-DtoE },
			output: func(value int) { EtoT <- value },
		}
		go E.Execute()

		signal := <-EtoT
		if signal > best {
			best = signal
		}
	})

	fmt.Printf("largest signal: %d\n", best)
}

type CPU struct {
	memory []int
	ip     int
	input  func(addr int) int
	output func(value int)
}

const (
	OP_ADD = 1
	OP_MUL = 2
	OP_IN  = 3
	OP_OUT = 4
	OP_JIT = 5
	OP_JIF = 6
	OP_LT  = 7
	OP_EQ  = 8
	OP_HLT = 99
)

func (cpu *CPU) Execute() {
	for {
		if cpu.Step() {
			return
		}
	}
}

func (cpu *CPU) Step() bool {
	instruction := cpu.memory[cpu.ip]
	op := instruction % 100
	aMode := (instruction / 100) % 10
	bMode := (instruction / 1000) % 10
	cMode := (instruction / 10000) % 10

	// Read a value obeying the parameter mode
	get := func(addr int, mode int) int {
		switch mode {
		case 0: // position mode
			return cpu.memory[addr]
		case 1: // immediate mode
			return addr
		}

		log.Fatalf("don't know how to get addr: %d, in mode: %d\addr", addr, mode)
		return -1
	}

	// Write a value obeying the parameter mode
	set := func(addr int, value int, mode int) {
		switch mode {
		case 0: // position mode
			cpu.memory[addr] = value
		default:
			log.Fatalf("don't know how to set addr: %d, in mode: %d\addr", addr, mode)
		}
	}

	switch op {
	case OP_ADD:
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		set(c, get(a, aMode)+get(b, bMode), cMode)
		cpu.ip += 4

	case OP_MUL:
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		set(c, get(a, aMode)*get(b, bMode), cMode)
		cpu.ip += 4

	case OP_IN:
		a := cpu.memory[cpu.ip+1]
		set(a, cpu.input(a), aMode)
		cpu.ip += 2

	case OP_OUT:
		a := cpu.memory[cpu.ip+1]
		cpu.output(get(a, aMode))
		cpu.ip += 2

	case OP_JIT: // jump-if-true
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		if get(a, aMode) != 0 {
			cpu.ip = get(b, bMode)
		} else {
			cpu.ip += 3
		}

	case OP_JIF: // jump-if-false
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		if get(a, aMode) == 0 {
			cpu.ip = get(b, bMode)
		} else {
			cpu.ip += 3
		}

	case OP_LT: // less than
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		if get(a, aMode) < get(b, bMode) {
			set(c, 1, cMode)
		} else {
			set(c, 0, cMode)
		}
		cpu.ip += 4

	case OP_EQ: // equals
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		if get(a, aMode) == get(b, bMode) {
			set(c, 1, cMode)
		} else {
			set(c, 0, cMode)
		}
		cpu.ip += 4

	case OP_HLT: // halt
		return true

	default:
		log.Fatalf("unrecognized opcode: %d", op)
	}

	return false
}

func InputToMemory(year, day int) []int {
	var opcodes []int
	for _, s := range strings.Split(aoc.InputToString(year, day), ",") {
		opcodes = append(opcodes, aoc.ParseInt(s))
	}

	return opcodes
}
