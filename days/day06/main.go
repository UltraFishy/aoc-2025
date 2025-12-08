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

type question struct {
	numbers  []int
	operator rune
	new      []int
}

func checkIfOperator(input string) bool {
	switch input {
	case "+", "*":
		return true
	default:
		return false
	}
}

func parseInput(input string) (ans []question) {

	split_rows := strings.Split(input, "\n")

	num_of_questions := len(strings.Split(split_rows[0], " "))
	ans = make([]question, num_of_questions)

	for _, line := range split_rows {

		question_arguments := strings.Fields(line)

		for i, arg := range question_arguments {

			if checkIfOperator(arg) {

				// Operator
				ans[i].operator = rune(arg[0])

			} else {

				// Numbers
				num, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("Error: %v\n", arg)
				}

				ans[i].numbers = append(ans[i].numbers, num)

			}
		}
	}

	return ans
}

// +==================================================+
// |                  PART SECTION                    |
// +==================================================+

func (q question) compute() int {

	if len(q.numbers) == 0 {
		return 0
	}

	total := q.numbers[0]

	switch q.operator {
	case '+':
		for i, n := range q.numbers {
			if i != 0 {
				total += n
			}
		}
	case '*':
		for i, n := range q.numbers {
			if i != 0 {
				total *= n
			}
		}
	}

	return total
}

func (q question) new_numbers() {

	if len(q.numbers) == 0 {
		return
	}

	// Get biggest num
	max_str := 0
	for _, num := range q.numbers {
		str_num := len(strconv.Itoa(num))
		if str_num > max_str {
			max_str = str_num
		}
	}

	// Loop backwards
	for i := max_str - 1; i > 0; i-- {
		// strconv.Itoa(num)
		// if len()
	}

}

func Part1(input string) int {
	questions := parseInput(input)

	total := 0

	for _, question := range questions {
		total += question.compute()
	}

	return total
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

	logger := utils.NewPartLogger("day06")

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
