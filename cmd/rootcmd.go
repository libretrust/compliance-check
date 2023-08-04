package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compliance-check",
	Short: "Compliance Check is a tool for checking compliance of used licenses in a project.",
	Long:  "Compliance Check is a tool for checking compliance of used licenses in a project. It allows denylisting and allowlisting of specific licenses.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
