package main

import (
	"appguard/internal/scanner"
	"fmt"
	"os"
)

func main() {
	fmt.Println("KEY:", os.Getenv("GEMINI_API_KEY"))

	if len(os.Args) < 2 {
		fmt.Println("Usage: astguard <repo-path>")
		return
	}

	root := os.Args[1]

	findings, err := scanner.ScanRepo(root)
	if err != nil {
		fmt.Println("Scan error:", err)
		return
	}

	for _, f := range findings {
		fmt.Printf("[%s] %s:%d\n%s\n\n",
			f.Type,
			f.File,
			f.Line,
			f.Code,
		)
	}
}
