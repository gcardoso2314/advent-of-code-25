package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gcardoso2314/advent-of-code-25/internal/util"
)

func partOne(lines []string) {
	operator_line := len(lines) - 1
	operators := []string{}
	numbers := [][]int{}
	for i, line := range lines {
		elements := strings.Split(line, " ")
		if i != operator_line {
			numbers = append(numbers, []int{})
			for _, numStr := range elements {
				if numStr == "" {
					continue
				}
				num, _ := strconv.Atoi(numStr)
				numbers[i] = append(numbers[i], num)
			}
		} else {
			for _, opStr := range elements {
				if opStr == "" {
					continue
				}
				operators = append(operators, opStr)
			}
		}
	}

	result := 0
	for i := range len(operators) {
		var problemResult int
		op := operators[i]
		switch op {
		case "+":
			problemResult = 0
			for _, numArray := range numbers {
				problemResult += numArray[i]
			}
		case "*":
			problemResult = 1
			for _, numArray := range numbers {
				problemResult *= numArray[i]
			}
		}
		result += problemResult
	}
	fmt.Printf("Solution to part one: %d\n", result)
}

func partTwo(lines []string) {
	result := 0
	lineLen := len(lines[0])
	currOperator := ""
	currNumbers := []int{}
	for i := range lineLen {
		col := lineLen - i - 1
		number := ""
		for _, line := range lines {
			switch line[col] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// append to current number
				number += string(line[col])
			case '*', '+':
				currOperator = string(line[col])
			}
		}
		if number == "" {
			// End of a problem
			problemResult := 0
			if currOperator == "*" {
				problemResult = 1
			}
			for _, num := range currNumbers {
				if currOperator == "*" {
					problemResult *= num
				} else {
					problemResult += num
				}
			}

			result += problemResult
			// reset
			currOperator = ""
			currNumbers = []int{}
		} else {
			num, _ := strconv.Atoi(number)
			currNumbers = append(currNumbers, num)
		}
	}
	if len(currNumbers) > 0 && currOperator != "" {
		// End of a problem
		problemResult := 0
		if currOperator == "*" {
			problemResult = 1
		}
		for _, num := range currNumbers {
			if currOperator == "*" {
				problemResult *= num
			} else {
				problemResult += num
			}
		}

		result += problemResult
	}

	fmt.Printf("Solution to part two: %d\n", result)
}

func main() {
	lines, _ := util.ReadLines(os.Args[1])

	partOne(lines)
	partTwo(lines)
}
