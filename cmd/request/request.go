package request

import (
	"errors"
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/internal"
)

var headersIncluded bool
var outputResponse bool
var bodyFn string

func init() {
	Cmd.Flags().BoolVarP(&headersIncluded, "include-headers", "i", false, "Read headers from the request body")
	Cmd.Flags().BoolVarP(&outputResponse, "output-response", "o", false, "Output response status line and headers")
	Cmd.Flags().StringVarP(&bodyFn, "body", "b", "", "Filename to read the request body from. Use - for stdin.")
}

var Cmd = &cobra.Command{
	Use:     "request [method] [url]",
	Short:   "Send a web request to a url and print the response",
	Aliases: []string{"get", "post", "put", "delete"},
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var method, url string

		if cmd.CalledAs() == "request" {
			if len(args) == 1 {
				method = "get"
				url = args[0]
			} else {
				method = args[0]
				url = args[1]
			}
		} else {
			if len(args) != 1 {
				cmd.Help()
				os.Exit(1)
			}

			method = cmd.CalledAs()
			url = args[0]
		}

		var input io.Reader
		var err error

		if bodyFn != "" {
			if bodyFn == "-" {
				if termutil.Isatty(os.Stdin.Fd()) {
					cobra.CheckErr(errors.New("Unable to read from stdin"))
				} else {
					input = os.Stdin
				}
			} else {
				input, err = os.Open(bodyFn)
				cobra.CheckErr(err)
			}
		}

		resp, err := internal.MakeRequest(method, url, input, headersIncluded)
		cobra.CheckErr(err)

		cobra.CheckErr(internal.PrintResponse(resp, outputResponse))
	},
}
