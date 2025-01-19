package please

import (
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/cmd/identify"
	"github.com/stilvoid/please/cmd/parse"
	"github.com/stilvoid/please/cmd/request"
	"github.com/stilvoid/please/cmd/respond"
	"github.com/stilvoid/please/cmd/serve"
)

var version = "git"

func init() {
	Cmd.Version = version

	Cmd.AddCommand(identify.Cmd)
	Cmd.AddCommand(parse.Cmd)
	Cmd.AddCommand(request.Cmd)
	Cmd.AddCommand(respond.Cmd)
	Cmd.AddCommand(serve.Cmd)

	Cmd.Root().CompletionOptions.DisableDefaultCmd = true
}

var Cmd = &cobra.Command{
	Use:     "please",
	Short:   "Please is a utility for making and receiving web requests and parsing and reformatting the common data formats that are sent over them.",
	Version: version,
}
