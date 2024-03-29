package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
	"strings"
)

func main() {
	var sum int
	for _, room := range InputToRooms() {
		if room.IsReal() {
			sum += room.SectorID
		}
	}

	fmt.Println(sum)
}

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

func (r Room) IsReal() bool {
	var counter aoc.FrequencyCounter[rune]
	for _, c := range r.Name {
		if c == '-' {
			continue
		}
		counter.Add(c)
	}

	entries := counter.Entries()
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count == entries[j].Count {
			return entries[i].Value < entries[j].Value
		}
		return entries[i].Count > entries[j].Count
	})

	var sb strings.Builder
	for _, entry := range entries[:5] {
		sb.WriteRune(entry.Value)
	}

	return sb.String() == r.Checksum
}

func InputToRooms() []Room {
	return aoc.InputLinesTo(2016, 4, func(line string) Room {
		hyphen := strings.LastIndex(line, "-")
		bracket := strings.LastIndex(line, "[")

		return Room{
			Name:     line[:hyphen],
			SectorID: aoc.ParseInt(line[hyphen+1 : bracket]),
			Checksum: line[bracket+1 : len(line)-1],
		}
	})
}
