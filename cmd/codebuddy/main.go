package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.1.0"

func main() {
	rootCmd := &cobra.Command{
		Use:   "codebuddy",
		Short: "CodeBuddy - Generate production-ready code from your terminal",
		Long: `CodeBuddy is an offline CLI tool that generates clean, ready-to-use code.
Pick a language, pick a pattern, answer a few questions — get code instantly.

No internet. No tokens. No AI costs. Just templates + your input = code.`,
	}

	rootCmd.AddCommand(
		newPythonCmd(),
		newVersionCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("codebuddy %s\n", Version)
		},
	}
}
