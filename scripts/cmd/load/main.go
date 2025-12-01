package main

import (
	"aoc-2025/utils"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const ENV = ".env"

func main() {

	// Load .env file
	if err := godotenv.Load(ENV); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Print the value of AOC_SESSION_COOKIE
	envVar := os.Getenv("AOC_SESSION_COOKIE")
	if envVar == "" {
		log.Println("AOC_SESSION_COOKIE is not set in the .env file")
		os.Exit(1)
	}

	utils.CopyToClipboard("export AOC_SESSION_COOKIE=" + envVar)
	fmt.Println("AOC_SESSION_COOKIE copied to clipboard")
	fmt.Println("Load by pasting and entering")
}
