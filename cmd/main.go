package main

import (
	"appguard/scanner"
	"fmt"
	"os"
)

func main() {

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
