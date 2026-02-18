package cmd

import (
	"appguard/internal/scanner"
	"fmt"

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

		for _, f := range findings {
			fmt.Printf("[%s] %s:%d\n%s\n",
				f.Type,
				f.File,
				f.Line,
				f.Code,
			)

			if useAI {
				fmt.Printf("AI Risk Explanation: ",
					f.LLMExplanation,
				)
			}

			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolVar(&useAI, "ai", false, "Enable AI explanations")
}
