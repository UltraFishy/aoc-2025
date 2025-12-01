package main

import (
	"flag"
	"time"

	skeleton "aoc-2025/scripts/templates"
)

func main() {
	today := time.Now()
	day := flag.Int("day", today.Day(), "day number to fetch, 1-25")
	flag.Parse()
	skeleton.Run(*day)
}
