// Package cmd wires the mental CLI using Cobra.
// Commands are thin — all logic lives in internal/.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mental",
	Short: "Cross-session memory and AI session manager",
	Long: `mental — manage memory across LLM sessions and search,
export, and manage AI assistant sessions
(opencode, claude, gemini, ...).

Docs: https://github.com/mrbrandao/mental`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
