package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/http"
)

var requestOptions http.Options

var requestCmd = &cobra.Command{
	Use:     "request <METHOD> <URL>",
	Aliases: []string{"get", "post", "put", "delete"},
	Short:   "Make a web request to URL using METHOD",
	Long: `Make a web request to URL using METHOD.
Data on stdin will be passed as the request body.
Request output will be written to stdout.

If you use one of the aliases for this command, METHOD is taken from the name of the alias.
For example: 'please get URL' is equivalent to 'please request GET URL'.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		method := cmd.CalledAs()

		if len(args) == 2 {
			method = args[0]
		} else if method == "request" {
			panic("You must supply both a method and a url")
		}

		url := args[len(args)-1]

		var input io.Reader
		if !termutil.Isatty(os.Stdin.Fd()) {
			input = os.Stdin
		}

		resp, err := http.MakeRequest(method, url, input, requestOptions)
		if err != nil {
			panic(fmt.Errorf("Unable to make request: %s", err.Error()))
		}

		respStr, err := http.FormatResponse(resp, requestOptions)
		if err != nil {
			panic(fmt.Errorf("Unable to output response: %s", err.Error()))
		}

		fmt.Println(respStr)
	},
}

func init() {
	requestCmd.Flags().BoolVarP(&requestOptions.HeadersIncluded, "included-headers", "i", false, "Flag that the input data already includes HTTP headers.")
	requestCmd.Flags().BoolVarP(&requestOptions.IncludeHeaders, "headers", "H", false, "Include HTTP headers in the output.")
	requestCmd.Flags().BoolVarP(&requestOptions.IncludeStatus, "status", "s", false, "Include the HTTP status code in the output.")

	Root.AddCommand(requestCmd)
}
