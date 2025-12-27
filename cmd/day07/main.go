package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

type Coord struct {
	row int
	col int
}

func NewCoord(row, col int) *Coord {
	return &Coord{row: row, col: col}
}

func findStart(grid [][]string) (*Coord, error) {
	for i, row := range grid {
		for j, val := range row {
			if val == "S" {
				return NewCoord(i, j), nil
			}
		}
	}

	return nil, errors.New("Could not find Start anywhere")
}

func iterateStep(grid [][]string, beams []*Coord, height, width int) ([]*Coord, int) {
	// Iterate through the beams downwards, return new beams and count of splits
	newBeams := []*Coord{}
	splitCount := 0
	visited := map[Coord]bool{}
	for _, beam := range beams {
		if visited[*beam] {
			continue
		}
		visited[*beam] = true

		// check if reached end of grid
		if beam.row+1 >= height {
			continue
		}

		// go down by 1
		newBeam := NewCoord(beam.row+1, beam.col)

		// check if split
		switch grid[newBeam.row][newBeam.col] {
		case ".":
			newBeams = append(newBeams, newBeam)
		case "^":
			splitCount += 1
			if newBeam.col-1 >= 0 {
				newBeams = append(newBeams, NewCoord(newBeam.row, newBeam.col-1))
			}
			if newBeam.col+1 < width {
				newBeams = append(newBeams, NewCoord(newBeam.row, newBeam.col+1))
			}
		default:
			log.Fatal("Unrecognized character in grid")
		}
	}

	return newBeams, splitCount
}

func partOne(grid [][]string, startPos *Coord) {
	height := len(grid)
	width := len(grid[0])
	beams := []*Coord{startPos}
	splitCount := 0
	for {
		var splits int
		if len(beams) == 0 {
			break
		}
		beams, splits = iterateStep(grid, beams, height, width)
		splitCount += splits
	}
	fmt.Printf("Solution to part one: %d\n", splitCount)
}

func countTimelines(pos Coord, grid [][]string, timelineCache map[Coord]int) int {
	// Check for cache and return
	count, ok := timelineCache[pos]
	if ok {
		return count
	}

	height := len(grid)
	width := len(grid[0])

	// check if reached end of grid
	if pos.row+1 >= height {
		timelineCache[pos] = 1
		return 1
	}

	// go down by 1
	newPos := NewCoord(pos.row+1, pos.col)

	// check if split
	switch grid[newPos.row][newPos.col] {
	case ".":
		count = countTimelines(*newPos, grid, timelineCache)
	case "^":
		var leftCount int
		var rightCount int
		if newPos.col-1 >= 0 {
			leftCount = countTimelines(*NewCoord(newPos.row, newPos.col-1), grid, timelineCache)
		}
		if newPos.col+1 < width {
			rightCount = countTimelines(*NewCoord(newPos.row, newPos.col+1), grid, timelineCache)
		}
		count = leftCount + rightCount

	default:
		log.Fatal("Unrecognized character in grid")
	}

	timelineCache[pos] = count

	return count
}

func partTwo(grid [][]string, startPos *Coord) {
	// depth first search and we cache the number of timelines
	timelineCache := map[Coord]int{}

	count := countTimelines(*startPos, grid, timelineCache)

	fmt.Printf("Solution to part two: %d\n", count)
}

func main() {
	lines, _ := util.ReadLines(os.Args[1])
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		grid[i] = strings.Split(line, "")
	}
	startPos, err := findStart(grid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(startPos)
	partOne(grid, startPos)
	partTwo(grid, startPos)
}
