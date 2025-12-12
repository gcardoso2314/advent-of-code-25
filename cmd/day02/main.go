package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

func isRepeatedN(str string, N int) bool {
	patterns := make(map[string]bool)
	for i := 0; i <= len(str)-N; i += N {
		patterns[str[i:i+N]] = true
	}

	return len(patterns) == 1
}

func partOne(line string) {
	result := 0

	ranges := strings.Split(line, ",")
	for _, rng := range ranges {
		rangeStartEnd := strings.Split(rng, "-")
		start, err := strconv.Atoi(rangeStartEnd[0])
		if err != nil {
			log.Fatal(err)
			return
		}
		end, err := strconv.Atoi(rangeStartEnd[1])
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := start; i <= end; i++ {
			// Determine if ID is invalid
			n_string := strconv.Itoa(i)
			len_str := len(n_string)

			if len_str%2 != 0 {
				continue
			} else if isRepeatedN(n_string, len_str/2) {
				result += i
			}
		}
	}

	fmt.Printf("Solution for part one: %d\n", result)
}

func partTwo(line string) {
	result := 0

	ranges := strings.Split(line, ",")
	for _, rng := range ranges {
		rangeStartEnd := strings.Split(rng, "-")
		start, err := strconv.Atoi(rangeStartEnd[0])
		if err != nil {
			log.Fatal(err)
			return
		}
		end, err := strconv.Atoi(rangeStartEnd[1])
		if err != nil {
			log.Fatal(err)
			return
		}
		for i := start; i <= end; i++ {
			// Determine if ID is invalid
			n_string := strconv.Itoa(i)
			len_str := len(n_string)

			// check factors up to sqrt(i)
			for f := 1; f < len_str; f++ {
				if len_str%f == 0 && isRepeatedN(n_string, f) {
					result += i
					break
				}
			}
		}
	}

	fmt.Printf("Solution for part two: %d\n", result)
}

func main() {
	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	// Only expecting one line
	if len(lines) != 1 {
		log.Fatal("Expected exactly one line")
		return
	}
	line := lines[0]

	partOne(line)
	partTwo(line)
}
