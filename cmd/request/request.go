package request

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stilvoid/please/internal"
)

var headersIncluded bool
var verbose bool
var inputFile string
var outputFile string
var headers []string

func init() {
	Cmd.Flags().BoolVarP(&headersIncluded, "include-headers", "i", false, "Read headers from the request body")
	Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Output response status line and headers")
	Cmd.Flags().StringVarP(&inputFile, "data", "d", "", "Filename to read the request body from. Omit for stdin.")
	Cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Filename to write the response to. Omit for stdout.")
	Cmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "Add a header to the request (can be used multiple times)")
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

		if inputFile == "" {
			input, err = internal.StdinOrNothing()
		} else {
			input, err = internal.FileOrStdin(inputFile)
		}
		cobra.CheckErr(err)

		// Convert header array to map
		headerMap := make(map[string][]string)
		for _, h := range headers {
			parts := strings.SplitN(h, ":", 2)
			if len(parts) != 2 {
				cobra.CheckErr("Invalid header format. Use 'Name: Value'")
			}
			name := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headerMap[name] = append(headerMap[name], value)
		}

		resp, err := internal.MakeRequest(method, url, input, headersIncluded, headerMap)
		cobra.CheckErr(err)

		cobra.CheckErr(internal.PrintResponse(resp, verbose, outputFile))
	},
}