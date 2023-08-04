package cmd

import (
	"fmt"

	"github.com/libretrust/compliance-check/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(analyzeCmd)
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Long:  "Analyze current repo for compliance.",
	Short: "Analyze current repo for compliance.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting analyzing the structure of the repository. This may take a while.")
		if len(args) == 0 {
			fmt.Println("No directories provided. Proceeding with scanning current working directory.")
			args = append(args, "./")
		}
		for _, path := range args {
			utils.RepositoryAnalyzer(path)
		}
	},
}
