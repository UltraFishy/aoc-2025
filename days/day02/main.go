package main

import (
	"aoc-2025/utils"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// +==================================================+
// |                 PARSE SECTION                    |
// +==================================================+

type IDRange struct {
	lower  int
	higher int
}

func parseInput(input string) (ans []IDRange) {

	// Note:
	// - Split by commas
	// - Split by -
	// - store lower and higher in a range struct
	// - output a list of ranges

	split_commas := strings.SplitSeq(input, ",")

	for ranges := range split_commas {
		split_dash := strings.Split(ranges, "-")

		lower, err := strconv.Atoi(split_dash[0])
		if err != nil {
			log.Fatal("[Error 1] String Conv to int error")
		}

		higher, err := strconv.Atoi(split_dash[1])
		if err != nil {
			log.Fatal("[Error 2] String Conv to int error")
		}

		new_range := IDRange{
			lower:  lower,
			higher: higher,
		}
		ans = append(ans, new_range)
	}

	return ans
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func checkForInvalid(value int) bool {

	// Note:
	// - convert to a string
	// - check if len of string is even
	// - if not then return false
	// - take the middle index
	// - loop through and compare the start and middle increment by i
	// - if anything doesn't match then return false
	// - otherwise return true

	string_value := strconv.Itoa(value)
	len_string := len(string_value)

	// Not Even
	if len_string%2 != 0 {
		return false
	}

	// Middle Index
	middle := (len_string / 2)

	for i := range middle {

		start_idx := i
		middle_idx := middle + i

		if string_value[start_idx] != string_value[middle_idx] {
			return false
		}

	}

	return true
}

func Part1(input string) int {

	// Note:
	// - loop through all ranges
	// - for a range start at the lower value
	// - loop until the higher value
	// - check each number for repeating
	// - if repeating add to total

	total := 0
	ranges := parseInput(input)

	for _, r := range ranges {

		for i := r.lower; i <= r.higher; i++ {

			if checkForInvalid(i) {
				total += i
			}

		}
	}

	return total
}

func checkForInvalid_v2(value int) bool {

	// Note:
	// create a window that slides over everything
	// if the window doesn't fit
	// have a temp window that takes account of all things the window has passed
	// then append the next non fitting number

	//   123123123
	//   [ 1 ]
	//   [ 1 2 ]
	//   [ 1 2 3 ]
	//   *tick*

	string_value := strconv.Itoa(value)
	window := []byte{string_value[0]}
	count := 0
	needed_count := len(string_value) - 1

	for i := 0; i < needed_count; i++ {

		// modulus of the index of the windows len
		if string_value[i] == window[i%len(window)] {
			count++
		} else {

			// Reset to start with new window
			window = append(window, string_value[i])
			i = 0
			count = 0
		}

		fmt.Printf("%v : %v\n", string_value[i], window)

	}

	if count == needed_count {
		return true
	}
	return false
}

func Part2(input string) int {

	total := 0
	ranges := parseInput(input)

	for _, r := range ranges {

		for i := r.lower; i <= r.higher; i++ {

			if checkForInvalid_v2(i) {
				total += i
			}

		}
	}

	return total

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

	logger := utils.NewPartLogger("day02")

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
