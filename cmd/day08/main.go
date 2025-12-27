package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

type JBPair struct {
	first    *JunctionBox
	second   *JunctionBox
	distance float64
}

func NewJBPair(first, second *JunctionBox) *JBPair {
	distance := first.Distance(second)
	return &JBPair{
		first,
		second,
		distance,
	}
}

type JunctionBox struct {
	x       float64
	y       float64
	z       float64
	circuit int
}

func NewJunctionBox(x, y, z float64, circuit int) *JunctionBox {
	return &JunctionBox{
		x,
		y,
		z,
		circuit,
	}
}

func (jb *JunctionBox) Distance(other *JunctionBox) float64 {
	return math.Pow(jb.x-other.x, 2.0) + math.Pow(jb.y-other.y, 2.0) + math.Pow(jb.z-other.z, 2.0)
}

func parseBoxes(lines []string) []*JunctionBox {
	junctionBoxes := make([]*JunctionBox, len(lines))
	for i, line := range lines {
		coords := []float64{}
		values := strings.Split(line, ",")
		for _, val := range values {
			num, _ := strconv.Atoi(val)
			coords = append(coords, float64(num))
		}
		box := NewJunctionBox(coords[0], coords[1], coords[2], i)
		junctionBoxes[i] = box
	}

	return junctionBoxes
}

func sortPairs(pairs []*JBPair) {
	slices.SortFunc(pairs, func(a, b *JBPair) int {
		if a.distance < b.distance {
			return -1
		} else if a.distance == b.distance {
			return 0
		} else {
			return 1
		}
	})
}

func partOne(junctionBoxes []*JunctionBox, pairs []*JBPair, nConnections int) {
	for i := range nConnections {
		minPair := pairs[i]
		// second pair will take number of first pair
		sourceCircuit := minPair.second.circuit
		targetCircuit := minPair.first.circuit
		for _, jb := range junctionBoxes {
			if jb.circuit == sourceCircuit {
				jb.circuit = targetCircuit
			}
		}
	}

	nBoxes := len(junctionBoxes)
	circuits := make([]int, nBoxes)
	for _, jb := range junctionBoxes {
		circuits[jb.circuit] += 1
	}

	slices.Sort(circuits)

	result := circuits[nBoxes-3] * circuits[nBoxes-2] * circuits[nBoxes-1]
	fmt.Printf("Solution for part one is: %d\n", result)
}

func partTwo(junctionBoxes []*JunctionBox, pairs []*JBPair) {
	nBoxes := len(junctionBoxes)

	var result float64
	for i := range len(pairs) {
		minPair := pairs[i]
		// second pair will take number of first pair
		sourceCircuit := minPair.second.circuit
		targetCircuit := minPair.first.circuit

		countTargetCircuit := 0
		for _, jb := range junctionBoxes {
			switch jb.circuit {
			case sourceCircuit:
				jb.circuit = targetCircuit
				countTargetCircuit++
			case targetCircuit:
				countTargetCircuit++
			}
		}

		if countTargetCircuit == nBoxes {
			// Connected all the boxes into one circuit
			result = minPair.second.x * minPair.first.x
			break
		}
	}

	fmt.Printf("Solution for part two is: %d\n", int(result))
}

func main() {
	lines, _ := util.ReadLines(os.Args[1])
	connectionsArg := os.Args[2]
	nConnections, _ := strconv.Atoi(connectionsArg)

	junctionBoxes := parseBoxes(lines)
	pairs := []*JBPair{}
	for i, box := range junctionBoxes {
		for j := i + 1; j < len(junctionBoxes); j++ {
			pair := NewJBPair(box, junctionBoxes[j])
			pairs = append(pairs, pair)
		}
	}

	sortPairs(pairs)

	partOne(junctionBoxes, pairs, nConnections)
	partTwo(junctionBoxes, pairs)
}
