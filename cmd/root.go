// Package cmd wires the ais CLI using Cobra.
// Commands are thin — all logic lives in internal/.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ais",
	Short: "AI session manager for multiple assistants",
	Long: `ais — search, export, and manage sessions
across AI assistants (opencode, claude, gemini, ...).

Docs: https://github.com/mrbrandao/ais`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
