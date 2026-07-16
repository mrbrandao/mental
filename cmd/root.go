// Package cmd wires the mental CLI using Cobra.
// Commands are thin — all logic lives in internal/.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mrbrandao/mental/cmd/session"
	"github.com/mrbrandao/mental/internal/config"
	"github.com/mrbrandao/mental/internal/extensions"
)

var rootCmd = &cobra.Command{
	Use:   "mental",
	Short: "AI Session Manager",
	Long: `mental — AI Session Manager

Search session history and manage session context memory across
sessions and providers (opencode, claude, gemini, ...).
Extensions are pluggable: built-in or external via MENTAL_DIR.

Docs: https://github.com/mrbrandao/mental`,
	// Register built-in extensions before any subcommand runs.
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("config: %w", err)
		}
		return extensions.RegisterBuiltins(cfg.Dir())
	},
}

func init() {
	rootCmd.AddCommand(session.Cmd)
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
