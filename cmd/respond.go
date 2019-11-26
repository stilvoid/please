package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
)

var respondHeadersIncluded bool
var respondIncludeMethod bool
var respondIncludePath bool
var respondIncludeHeaders bool

var respondCmd = &cobra.Command{
	Use:   "respond <STATUS> [[<ADDRESS>:]<PORT>]",
	Short: "Listen for an HTTP request and respond to it.",
	Long: `Listens on the specified ADDRESS and PORT and responds with the chosen STATUS code.
Any data on stdin will be used as the body of the response.
The request body will be printed to stdout.

If you do not supply PORT, it defaults to 8000.
If you do not supply AdDRESS, it defaults to 0.0.0.0.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		status, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Errorf("Invalid status code: %s", args[0]))
		}

		address := "0.0.0.0:8000"
		if len(args) > 1 {
			address = args[1]
		}

		handler := responder{
			status:          status,
			includeHeaders:  respondIncludeHeaders,
			includeMethod:   respondIncludeMethod,
			includeUrl:      respondIncludePath,
			headersIncluded: respondHeadersIncluded,
		}

		listener, err := net.Listen("tcp", address)
		if err != nil {
			panic(fmt.Errorf("Unable to start HTTP listener: %s", err.Error()))
		}

		handler.listener = listener

		if termutil.Isatty(os.Stdin.Fd()) {
			handler.data = nil
		} else {
			handler.data = os.Stdin
		}

		server := &http.Server{Addr: address, Handler: handler}

		server.Serve(listener)
	},
}

func init() {
	respondCmd.Flags().BoolVarP(&respondHeadersIncluded, "included-headers", "i", false, "Flag that the input data includes HTTP headers.")
	respondCmd.Flags().BoolVarP(&respondIncludeMethod, "method", "m", false, "Include the HTTP request method in the output.")
	respondCmd.Flags().BoolVarP(&respondIncludePath, "path", "p", false, "Include the requested path in the output.")
	respondCmd.Flags().BoolVarP(&respondIncludeHeaders, "headers", "H", false, "Include HTTP request headers in the output.")

	Root.AddCommand(respondCmd)
}
