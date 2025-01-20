package main

import (
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/cmd/please"
)

func main() {
	cobra.CheckErr(please.Cmd.Execute())
}
