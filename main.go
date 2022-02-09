package main

import (
	"github.com/khodemobin/pilo/auth/cmd"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{
		Use:                "app",
		DisableAutoGenTag:  true,
		DisableSuggestions: true,
		Run: func(c *cobra.Command, args []string) {
			cmd.Execute()
		},
	}
	rootCmd.AddCommand(cmd.ServeCommand(), cmd.MigrateCommand(), cmd.SeedCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
