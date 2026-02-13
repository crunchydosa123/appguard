package main

import (
	"appguard/parser"
	"appguard/rules"
	"appguard/scanner"
	"fmt"
	"os"
)

func main() {

	code, _ := os.ReadFile("test.js")

	tree, _ := parser.Parse(code)

	var findings []rules.Finding

	scanner.Walk(tree.RootNode(), code, &findings)

	for _, f := range findings {
		fmt.Println("Trigger:", f.Type)
		fmt.Println("Code:", f.Code)
	}
}
