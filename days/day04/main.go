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

func parseInput(input string) (ans [][]int) {

	lines := strings.SplitSeq(input, "\n")
	grid := [][]int{}

	for line := range lines {
		cells := strings.SplitSeq(line, "")
		cells_int := []int{}

		for cell := range cells {
			if cell == "@" {
				cells_int = append(cells_int, 1)
			} else {
				cells_int = append(cells_int, 0)
			}
		}

		grid = append(grid, cells_int)
	}

	return grid
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

type direction struct {
	vertical   int
	horizontal int
}

func numAdjacent(target_x int, target_y int, grid [][]int) int {

	adjacent_positions := [8]direction{
		{vertical: -1, horizontal: -1},
		{vertical: -1, horizontal: 0},
		{vertical: -1, horizontal: 1},
		{vertical: 0, horizontal: -1},
		{vertical: 0, horizontal: 1},
		{vertical: 1, horizontal: -1},
		{vertical: 1, horizontal: 0},
		{vertical: 1, horizontal: 1},
	}

	count := 0

	for _, pos := range adjacent_positions {

		x := target_x + pos.horizontal
		y := target_y + pos.vertical

		if x < 0 || x > len(grid[0])-1 {
			continue
		}

		if y < 0 || y > len(grid)-1 {
			continue
		}

		if grid[y][x] == 1 {
			count++
		}

	}

	return count

}

func Part1(input string) int {
	grid := parseInput(input)
	count := 0

	for i, rows := range grid {
		for j := range rows {

			roll_count := numAdjacent(i, j, grid)
			if roll_count < 4 && grid[j][i] == 1 {
				count++
			}
		}
	}

	return count
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

	logger := utils.NewPartLogger("day04")

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
