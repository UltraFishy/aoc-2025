package main

import (
	"aoc-2025/utils"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// +==================================================+
// |                 PARSE SECTION                    |
// +==================================================+

type Instruction struct {
	Direction rune
	Distance  int
}

func parseInput(input string) (ans []Instruction) {
	// Into lines
	lines := strings.Split(input, "\n")
	instructions := []Instruction{}

	for _, val := range lines {

		// Extract letter from distance
		direction := rune(val[0])
		distance, err := strconv.Atoi(val[1:])

		// Error parsing
		if err != nil {
			log.Printf("Error parsing int")
			os.Exit(1)
		}

		// Into instruction
		instruction := Instruction{
			Direction: direction,
			Distance:  distance,
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func Part1(input string) int {
	instructions := parseInput(input)

	// Starts at 50
	current := 50
	zeros := 0

	for _, instruction := range instructions {

		if instruction.Direction == rune('L') {
			current = current - instruction.Distance
		} else {
			current = current + instruction.Distance
		}

		current = current % 100

		if current < 0 {
			current = 100 + current
		}

		if current == 0 {
			zeros++
		}

	}

	return zeros
}

func Part2(input string) int {
	instructions := parseInput(input)

	zero_count := 0
	current_index := 50

	// Array of 0's with index 0 having value of 1
	var arr [100]int
	arr[0] = 1

	// Loop over instructions
	for _, instruction := range instructions {

		// Loop for the distance
		for range instruction.Distance {

			// Apply dir
			if instruction.Direction == rune('L') {
				current_index--
			} else {
				current_index++
			}

			// Loop condition
			current_index = current_index % 100
			if current_index < 0 {
				current_index = 100 + current_index
			}

			// Check for zero
			if arr[current_index] == 1 {
				zero_count++
			}

		}

	}

	return zero_count
}

// +==================================================+
// |                  MAIN SECTION                    |
// +==================================================+

//go:embed input.txt
var input string

func init() {
	// for lines
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	logger := utils.NewPartLogger("day 1")

	switch part {
	case 1:
		ans := Part1(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		logger.PrintPart(1, ans)
	case 2:
		ans := Part2(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		logger.PrintPart(2, ans)
	default:
		ans1 := Part1(input)
		ans2 := Part2(input)
		logger.PrintPart(1, ans1)
		logger.PrintPart(2, ans2)
	}
}
