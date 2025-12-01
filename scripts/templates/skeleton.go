// Package skeleton makes skeletons to be filled out with solutions.
package skeleton

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"aoc-2025/utils"
)

//go:embed tmpls/*.go
var fs embed.FS

// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	ts, err := template.ParseFS(fs, "tmpls/*.go")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	mainFilename := filepath.Join(utils.Dirname(), "../../", fmt.Sprintf("days/day%02d/main.go", day))
	testFilename := filepath.Join(utils.Dirname(), "../../", fmt.Sprintf("days/day%02d/main_test.go", day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}
	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	ts.ExecuteTemplate(testFile, "main_test.go", nil)
	fmt.Printf("templates made for day%d\n", day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
