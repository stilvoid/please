package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use:  "please",
	Long: "Please is a utility for interacting with web APIs and common data formats",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() {
	if err := Root.Execute(); err != nil {
		os.Exit(1)
	}
}
