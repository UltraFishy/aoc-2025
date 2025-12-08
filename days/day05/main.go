package main

import (
	"aoc-2025/utils"
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// +==================================================+
// |                 PARSE SECTION                    |
// +==================================================+

type IRange struct {
	lower  int
	higher int
}

func parseInput(input string) ([]IRange, []int) {

	// NOTES:
	// - split to ranges and ingredients
	// - type for ranges []IRange
	// - ingredients []int

	split_range_ingre := strings.Split(input, "\n\n")

	split_range := strings.SplitSeq(split_range_ingre[0], "\n")
	split_ingre := strings.SplitSeq(split_range_ingre[1], "\n")

	ranges := []IRange{}

	for r := range split_range {
		split_r := strings.Split(r, "-")

		lower, err := strconv.Atoi(split_r[0])
		if err != nil {
			fmt.Println("Error: int conv 1")
		}

		higher, err := strconv.Atoi(split_r[1])
		if err != nil {
			fmt.Println("Error: int conv 2")
		}

		current_range := IRange{
			lower:  lower,
			higher: higher,
		}

		ranges = append(ranges, current_range)
	}

	ingredients := []int{}

	for i := range split_ingre {

		i_int, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Error: int conv 3")
		}

		ingredients = append(ingredients, i_int)
	}

	return ranges, ingredients
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+
func Part1(input string) int {
	ranges, ingredients := parseInput(input)
	count := 0

	for _, ing := range ingredients {
		for _, r := range ranges {
			if ing >= r.lower && ing <= r.higher {
				count++
				break
			}
		}
	}

	return count
}

func Part2(input string) int {
	ranges, _ := parseInput(input)

	// Sort ranges by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].lower < ranges[j].lower
	})

	// Merge
	merged := []IRange{}
	current := ranges[0]

	for _, r := range ranges[1:] {
		if r.lower <= current.higher+1 {

			// Overlapping or touching -> mergeee
			if r.higher > current.higher {
				current.higher = r.higher
			}
		} else {

			// No overlap -> push current
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)

	count := 0
	for _, r := range merged {
		count += (r.higher - r.lower) + 1
	}

	return count
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

	logger := utils.NewPartLogger("day05")

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
