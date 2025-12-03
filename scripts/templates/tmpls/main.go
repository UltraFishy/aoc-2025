package main

import (
	"aoc-2025/utils"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

// +==================================================+
// |                 PARSE SECTION                    |
// +==================================================+

func parseInput(input string) (ans []int) {
	_ = input
	ans = []int{0, 1, 2}
	return ans
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func Part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 1
}

func Part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

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

	logger := utils.NewPartLogger("")

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
