package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

func mod(x, y int) int {
	x = x % y
	if x >= 0 {
		return x
	}
	return y + x
}

func partOne() {
	pos := 50
	base := 100
	password := 0

	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range lines {
		dir := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
			return
		}
		switch dir {
		case 'L':
			pos -= distance
		case 'R':
			pos += distance
		}
		pos = mod(pos, base)
		if pos == 0 {
			password += 1
		}
	}

	fmt.Printf("Password is %d\n", password)
}

func partTwo() {
	pos := 50
	base := 100
	password := 0

	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range lines {
		dir := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
			return
		}
		// Need to account for complete turns of the wheel
		numTurns := distance / base
		password += numTurns

		// update distance to only do the remainder of turn
		distance -= numTurns * base

		new_pos := pos
		switch dir {
		case 'L':
			new_pos -= distance
		case 'R':
			new_pos += distance
		}
		mod_pos := mod(new_pos, base)

		if mod_pos == 0 || (mod_pos != new_pos && pos != 0) {
			password += 1
		}

		pos = mod_pos

	}

	fmt.Printf("Password is %d\n", password)
}

func main() {
	partOne()
	partTwo()
}
