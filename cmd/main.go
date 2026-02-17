package main

import (
	"appguard/internal/llm"
	"appguard/internal/scanner"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: astguard <repo-path>")
		return
	}

	root := os.Args[1]

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env not loaded")
	}

	findings, err := scanner.ScanRepo(root)
	if err != nil {
		fmt.Println("Scan error:", err)
		return
	}

	ctx := context.Background()

	client, err := llm.NewClient(ctx)
	if err != nil {
		fmt.Println("LLM client error:", err)
	}

	enrichedFindings, err := llm.EnrichFindings(ctx, client, findings)
	if err != nil {
		fmt.Println("LLM enrichment failed, continuing with raw findings")
		enrichedFindings = findings
	}

	for _, f := range enrichedFindings {
		fmt.Printf("[%s] %s:%d\n%s\n%s\n\n",
			f.Type,
			f.File,
			f.Line,
			f.Code,
			f.LLMExplanation,
		)
	}
}
