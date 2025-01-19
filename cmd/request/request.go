package request

import (
	"io"
	"log"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/internal/web"
)

var headersIncluded bool
var outputHeaders bool
var outputStatus bool
var bodyFn string

func init() {
	Cmd.Flags().BoolVarP(&headersIncluded, "include-headers", "i", false, "Read headers from the request body")
	Cmd.Flags().BoolVarP(&outputHeaders, "output-headers", "o", false, "Output response headers")
	Cmd.Flags().BoolVarP(&outputStatus, "output-status", "s", false, "Output response status line")
	Cmd.Flags().StringVarP(&bodyFn, "body", "b", "", "Filename to read the request body from. Use - for stdin.")
}

var Cmd = &cobra.Command{
	Use:     "request [method] [url]",
	Short:   "Make a web request to [url] using the specified HTTP [method].",
	Aliases: []string{"get", "post", "put", "delete"},
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var method, url string

		if cmd.CalledAs() == "request" {
			if len(args) != 2 {
				cmd.Help()
				os.Exit(1)
			}

			method = args[0]
			url = args[1]
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
					log.Fatal("Unable to read from stdin")
				} else {
					input = os.Stdin
				}
			} else {
				input, err = os.Open(bodyFn)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		resp, err := web.MakeRequest(method, url, input, headersIncluded)
		if err != nil {
			log.Fatal(err)
		}

		err = web.WriteResponse(os.Stdout, resp, outputHeaders, outputStatus)
		if err != nil {
			log.Fatal(err)
		}
	},
}
