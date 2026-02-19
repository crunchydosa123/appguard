package cmd

import (
	"appguard/internal/llm"
	"appguard/internal/scanner"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var useAI bool

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "scan repository for security risks",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		findings, err := scanner.ScanRepo(path)
		if err != nil {
			fmt.Println("Scan error: ", err)
			return
		}

		if useAI && len(findings) > 0 {
			client, err := llm.NewClient(cmd.Context())
			if err != nil {
				fmt.Println("LLM client error:", err)
				return
			}

			findings, err = llm.EnrichFindings(cmd.Context(), client, findings)
			if err != nil {
				fmt.Println("LLM enrichment error:", err)
				return
			}
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{"Type", "File", "Line", "LLM Explanation"})

		for _, f := range findings {
			t.AppendRow(table.Row{f.Type, f.File, f.Line, f.LLMExplanation})
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolVar(&useAI, "ai", false, "Enable AI explanations")
}
