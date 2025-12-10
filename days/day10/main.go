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

type Button struct {
	lights []bool
}

type Joltage struct {
	joltages []int
}

type Lights struct {
	lights []bool
}

type Machine struct {
	lights   Lights
	buttons  []Button
	joltages Joltage
}

func parseInput(input string) (ans []Machine) {
	split_lines := strings.SplitSeq(input, "\n")

	for line := range split_lines {

		split_whitespace := strings.Fields(line)
		last := len(split_whitespace) - 1

		// ====================
		//        Lights
		// ====================
		lights_string := strings.Split(strings.Trim(split_whitespace[0], "[]"), "")
		largest_light := len(lights_string)

		lights := Lights{[]bool{}}

		// Lights to bitmap
		for _, string := range lights_string {
			switch string {
			case ".":
				lights.lights = append(lights.lights, false)
			default:
				lights.lights = append(lights.lights, true)
			}
		}

		// ====================
		//       Buttons
		// ====================
		buttons_string := split_whitespace[1 : last-1]
		buttons := []Button{}

		for _, butts := range buttons_string {

			buttons_ := strings.Split(strings.Trim(butts, "()"), ",")
			button := Button{make([]bool, largest_light)}

			// To Bitmap
			for _, b := range buttons_ {
				b_, err := strconv.Atoi(b)
				if err != nil {
					fmt.Printf("Error: %v\n", b)
				}

				button.lights[b_] = true
			}

			buttons = append(buttons, button)
		}

		// ====================
		//       Buttons
		// ====================

		joltages_string := strings.Split(strings.Trim(split_whitespace[last], "{}"), ",")
		joltages := Joltage{[]int{}}

		for _, j := range joltages_string {
			j_, err := strconv.Atoi(j)
			if err != nil {
				fmt.Printf("Error: %v\n", j)
			}

			joltages.joltages = append(joltages.joltages, j_)
		}

		machine := Machine{
			lights:   lights,
			buttons:  buttons,
			joltages: joltages,
		}

		ans = append(ans, machine)
	}

	return ans
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func Part1(input string) int {
	machines := parseInput(input)

	fmt.Println(machines)

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

	logger := utils.NewPartLogger("day10")

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
