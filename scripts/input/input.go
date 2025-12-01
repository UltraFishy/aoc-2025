package input

import (
	"fmt"
	"path/filepath"
	"strings"

	"aoc-2025/utils"
)

func GetInput(day int, cookie string) {
	fmt.Printf("fetching for day %d, year %d\n", day, 2025)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", 2025, day)
	body := GetWithAOCCookie(url, cookie)

	if strings.HasPrefix(string(body), "Puzzle inputs differ by user") {
		panic("'Puzzle inputs differ by user' response")
	}

	// write to file
	filename := filepath.Join(utils.Dirname(), "../..", fmt.Sprintf("days/day%02d/input.txt", day))
	WriteToFile(filename, body)

	fmt.Println("Wrote to file: ", filename)

	fmt.Println("Done!")
}
