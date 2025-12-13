package main

import (
	"fmt"
	"os"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

func parseGrid(lines []string) [][]bool {
	grid := make([][]bool, len(lines))
	for i := range lines {
		grid[i] = make([]bool, len(lines[i]))
		for j, b := range lines[i] {
			grid[i][j] = b == '@'
		}
	}
	return grid
}

func partOne(grid [][]bool) {
	result := 0
	nRows := len(grid)
	nCols := len(grid[0])

	directions := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	for row := range nRows {
		for col := range nCols {
			if grid[row][col] {
				//check surroundings
				countAdjRolls := 0
				for _, dir := range directions {
					nrow := row + dir[0]
					ncol := col + dir[1]
					if nrow >= 0 && nrow < nRows && ncol >= 0 && ncol < nCols && grid[nrow][ncol] {
						countAdjRolls++
					}
				}
				if countAdjRolls < 4 {
					result++
				}
			}
		}
	}

	fmt.Printf("Solution for part one: %d\n", result)
}

func partTwo(grid [][]bool) {
	result := 0
	nRows := len(grid)
	nCols := len(grid[0])

	directions := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	rollsRemoved := [][]int{}

	for {
		for row := range nRows {
			for col := range nCols {
				if grid[row][col] {
					//check surroundings
					countAdjRolls := 0
					for _, dir := range directions {
						nrow := row + dir[0]
						ncol := col + dir[1]
						if nrow >= 0 && nrow < nRows && ncol >= 0 && ncol < nCols && grid[nrow][ncol] {
							countAdjRolls++
						}
					}
					if countAdjRolls < 4 {
						result++
						rollsRemoved = append(rollsRemoved, []int{row, col})
					}
				}
			}
		}
		if len(rollsRemoved) == 0 {
			break
		}
		for _, coord := range rollsRemoved {
			grid[coord[0]][coord[1]] = false
		}
		rollsRemoved = [][]int{}
	}

	fmt.Printf("Solution for part two: %d\n", result)

}

func main() {
	lines, _ := util.ReadLines(os.Args[1])
	grid := parseGrid(lines)
	partOne(grid)
	partTwo(grid)
}
