package main

import (
	"aoc-2025/utils"
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// +==================================================+
// |                 PARSE SECTION                    |
// +==================================================+

func parseInput(input string) (ans [][]int) {
	split_line := strings.SplitSeq(input, "\n")
	int_line := [][]int{}

	for line := range split_line {
		split_batteries := strings.SplitSeq(line, "")
		int_batteries := []int{}

		for battery := range split_batteries {
			batt, err := strconv.Atoi(battery)
			if err != nil {
				fmt.Printf("Error parsing int")
			}

			int_batteries = append(int_batteries, batt)
		}

		int_line = append(int_line, int_batteries)
	}

	return int_line
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func Part1(input string) int {
	batteries := parseInput(input)

	// Notes:
	// - loop through a battery pack
	// - save highest battery and second highest
	// - if new highest battery replace old highest

	joltage_total := 0

	for _, bank := range batteries {

		current_highest_one := 0
		current_highest_two := 0

		for i, batt := range bank {

			if batt > current_highest_one && i != len(bank)-1 {
				current_highest_one = batt
				current_highest_two = 0
				continue
			}

			if batt > current_highest_two {
				current_highest_two = batt
			}

		}

		// Join the highest together
		joltage_string := fmt.Sprintf("%v%v", current_highest_one, current_highest_two)
		joltage, err := strconv.Atoi(joltage_string)
		if err != nil {
			fmt.Println("Error conv int")
		}

		joltage_total += joltage
	}

	return joltage_total
}

func Part2(input string) int {
	batteries := parseInput(input)

	for _, bank := range batteries {

		joltage_array := [12]int{}

		// Notes:
		// - loop backwards through bank
		//    - loop through joltage array
		//    - check if current joltage is larger than index
		//    - next index
		//    - if equal to last index, check second last and so on

		fmt.Println(joltage_array)

		j := 12 - 1

		for i := len(bank) - 1; i > 0; i-- {

			if j == 0 {
				if bank[i] == joltage_array[j] {
					j++
					i++
					continue
				}
			}

			// Fills the joltage bank
			if bank[i] > joltage_array[j] {
				joltage_array[j] = bank[i]
				j--
			}

			// Makes sure it doesn't break
			if j < 0 {
				j = 0
			}

			fmt.Println(joltage_array)

		}

	}

	return 2
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

	logger := utils.NewPartLogger("day03")

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
