package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

func stringToIntArray(s string) []int {
	intArray := make([]int, len(s))

	for i, digitChar := range s {
		intArray[i] = int(digitChar - '0')
	}

	return intArray
}

func argMax(a []int) int {
	// Return the minimum index that matches the largest value
	index := 0
	maxVal := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > maxVal {
			index = i
			maxVal = a[i]
		}
	}
	return index
}

func findLargestJoltage(line string, N int) string {
	if N == 0 {
		return ""
	}
	// Find the max digit in 0:len(line)-N+1
	// Then recurse to N-1 for index:len(line)-1
	digits := stringToIntArray(line[:len(line)-N+1])
	maxIndex := argMax(digits)

	return string(line[maxIndex:maxIndex+1]) + findLargestJoltage(line[maxIndex+1:], N-1)
}

func partOne(lines []string) {
	result := 0
	for _, line := range lines {
		strResult := findLargestJoltage(line, 2)
		lineResult, _ := strconv.Atoi(strResult)
		result += lineResult
	}

	fmt.Printf("Solution for part one: %d\n", result)
}

func partTwo(lines []string) {
	result := 0
	for _, line := range lines {
		strResult := findLargestJoltage(line, 12)
		lineResult, _ := strconv.Atoi(strResult)
		result += lineResult
	}

	fmt.Printf("Solution for part one: %d\n", result)
}

func main() {
	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	partOne(lines)
	partTwo(lines)
}
