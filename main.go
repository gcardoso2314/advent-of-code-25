package main

import (
	"fmt"
	"log"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
