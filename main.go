package main

import (
	"github.com/khodemobin/pilo/auth/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:                "app",
		DisableAutoGenTag:  true,
		DisableSuggestions: true,
	}
	rootCmd.AddCommand(cmd.ServeCommand(), cmd.MigrateCommand())
	rootCmd.Execute()
}
