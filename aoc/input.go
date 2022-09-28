package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// InputFilename determines the input file for a specific day's part.
func InputFilename(year, day int) string {
	return fmt.Sprintf("cmd/%d/%02d-1/input.txt", year, day)
}

// InputToBytes reads the entire input file into a slice of bytes.
func InputToBytes(year, day int) []byte {
	bs, err := ioutil.ReadFile(InputFilename(year, day))
	if err != nil {
		log.Fatalf("unable to read input.txt: %+v", err)
	}

	return bytes.TrimSpace(bs)
}

// InputToString reads the entire input file into a string.
func InputToString(year, day int) string {
	return string(InputToBytes(year, day))
}

// InputToLines reads the input file into a slice of strings with each string
// representing a line of the file.  The newline character is not included.
func InputToLines(year, day int) []string {
	file, err := os.Open(InputFilename(year, day))
	if err != nil {
		log.Fatalf("unable to open input.txt: %+v", err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading input.txt: %+v", err)
	}

	return lines
}

// InputLinesTo transforms each line of the input file into an instance
// returned by a transform function.  The instances are returned in a
// slice in the same order as they appear in the file.
func InputLinesTo[T any](year, day int, parse func(string) (T, error)) []T {
	var ts []T
	for _, line := range InputToLines(year, day) {
		t, err := parse(line)
		if err != nil {
			log.Fatalf("unable to parse line '%s': %+v", line, err)
		}

		ts = append(ts, t)
	}

	return ts
}

// InputToInt reads the input file into a single integer.
func InputToInt(year, day int) int {
	return InputToInts(year, day)[0]
}

// InputToInts reads the input file into a slice of integers.
func InputToInts(year, day int) []int {
	return InputLinesTo(year, day, strconv.Atoi)
}

// InputToGrid2D builds a Grid2D instance from the input using the provided
// function to determine the value for each cell in the grid.
func InputToGrid2D[T any](year, day int, fn func(int, int, string) T) Grid2D[T] {
	lines := InputToLines(year, day)

	grid := NewGrid2D[T](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			grid.AddXY(x, y, fn(x, y, string(c)))
		}
	}

	return grid
}
