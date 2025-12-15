package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

type IngredientRange struct {
	low  int
	high int
}

func partOne(ingredientRanges []IngredientRange, availableIngredients []int) {
	result := 0
	for _, ingred := range availableIngredients {
		for _, iRange := range ingredientRanges {
			if ingred >= iRange.low && ingred <= iRange.high {
				result++
				break
			}
		}
	}

	fmt.Printf("Solution for part one: %d\n", result)
}

// overlaps returns true if the two ranges overlap
func overlaps(a, b IngredientRange) bool {
	return a.low <= b.high && b.low <= a.high
}

// contains returns true if range a completely contains range b
func contains(a, b IngredientRange) bool {
	return a.low <= b.low && b.high <= a.high
}

func partTwo(ingredientRanges []IngredientRange) {
	outputRanges := []IngredientRange{}

	for _, newRange := range ingredientRanges {
		// Queue of ranges to process, starting with the new range
		queue := []IngredientRange{newRange}

		for len(queue) > 0 {
			// Process first range in queue
			current := queue[0]
			queue = queue[1:]

			// Check if this range is entirely contained in any existing range
			fullyContained := false
			for _, existing := range outputRanges {
				if contains(existing, current) {
					fullyContained = true
					break
				}
			}
			if fullyContained {
				continue // Skip this range, it's already covered
			}

			// Check for overlaps and split if needed
			overlapped := false
			for _, existing := range outputRanges {
				if !overlaps(current, existing) {
					continue
				}

				// Split the current range around the existing one
				overlapped = true
				if current.low < existing.low {
					queue = append(queue, IngredientRange{low: current.low, high: existing.low - 1})
				}
				if current.high > existing.high {
					queue = append(queue, IngredientRange{low: existing.high + 1, high: current.high})
				}
				break // Only need to split once per iteration
			}

			// If no overlaps, add to output
			if !overlapped {
				outputRanges = append(outputRanges, current)
			}
		}
	}

	// Sum up all the unique non-overlapping ranges
	result := 0
	for _, r := range outputRanges {
		result += r.high - r.low + 1
	}

	fmt.Printf("Solution for part two: %d\n", result)
}

func parseLines(lines []string) ([]IngredientRange, []int) {
	ingredientRanges := []IngredientRange{}
	availableIngredients := []int{}
	idx := 0
	for _, line := range lines {
		idx++
		if line == "" {
			break
		}
		iRange := strings.Split(line, "-")
		low, _ := strconv.Atoi(iRange[0])
		high, _ := strconv.Atoi(iRange[1])
		ingredientRanges = append(ingredientRanges, IngredientRange{low, high})
	}
	for i := idx; i < len(lines); i++ {
		numStr := lines[i]
		num, _ := strconv.Atoi(numStr)
		availableIngredients = append(availableIngredients, num)
	}

	return ingredientRanges, availableIngredients
}

func main() {
	lines, _ := util.ReadLines(os.Args[1])
	ingredientRanges, availableIngredients := parseLines(lines)
	partOne(ingredientRanges, availableIngredients)
	partTwo(ingredientRanges)
}
