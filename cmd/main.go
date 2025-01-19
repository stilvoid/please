package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/stilvoid/please/cmd/identify"
	"github.com/stilvoid/please/cmd/parse"
	"github.com/stilvoid/please/cmd/request"
	"github.com/stilvoid/please/cmd/respond"
	"github.com/stilvoid/please/cmd/serve"
)

var version = "git"

func init() {
	rootCmd.Version = version

	rootCmd.AddCommand(identify.Cmd)
	rootCmd.AddCommand(parse.Cmd)
	rootCmd.AddCommand(request.Cmd)
	rootCmd.AddCommand(respond.Cmd)
	rootCmd.AddCommand(serve.Cmd)

	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	log.SetFlags(0)
}

var rootCmd = &cobra.Command{
	Use:     "please",
	Short:   "Please is a utility for making and receiving web requests and parsing and reformatting the common data formats that are sent over them.",
	Version: version,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
