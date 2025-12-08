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

func kmpPrefixFunction(input string) []int {
	n := len(input)
	pi := make([]int, n)

	for i := 1; i < n; i++ {
		j := pi[i-1]

		for j > 0 && input[i] != input[j] {
			j = pi[j-1]
		}

		if input[i] == input[j] {
			j++
		}

		pi[i] = j
	}

	return pi
}

func checkForInvalid_v2(value int) bool {

	// Note:
	// Uses KMP algorithm
	// Count starting zeros,
	// mod with last index
	// if mod == 0 then patterns fit

	string_value := strconv.Itoa(value)
	result := kmpPrefixFunction(string_value)

	count := 0

	for i, num := range result {

		if num == 0 {
			if i > 0 {
				count += result[i-1] + 1
			} else {
				count++
			}
			continue
		}

	}

	last_index := len(result) - 1

	if result[last_index] != 0 && result[last_index]%count == 0 {
		return true
	}

	return false
}

func Part2(input string) int {

	total := 0
	ranges := parseInput(input)
	saved := []int{}

	for _, r := range ranges {

		for i := r.lower; i <= r.higher; i++ {

			if checkForInvalid_v2(i) {
				total += i
				saved = append(saved, i)

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
