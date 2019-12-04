package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
	httplease "github.com/stilvoid/please/http"
)

var respondOptions httplease.Options

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

		handler := httplease.Responder{
			Status:  status,
			Options: respondOptions,
			Stop:    make(chan bool),
		}

		if termutil.Isatty(os.Stdin.Fd()) {
			// Don't try to read from the terminal
			handler.Data = nil
		} else {
			handler.Data = os.Stdin
		}

		server := &http.Server{
			Addr:    address,
			Handler: handler,
		}

		go func() {
			<-handler.Stop
			server.Shutdown(context.TODO())
		}()

		fmt.Printf("Listening on %s...\n", address)
		if err = server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	},
}

func init() {
	respondCmd.Flags().BoolVarP(&respondOptions.HeadersIncluded, "included-headers", "i", false, "Flag that the input data includes HTTP headers.")
	respondCmd.Flags().BoolVarP(&respondOptions.IncludeMethod, "method", "m", false, "Include the HTTP request method in the output.")
	respondCmd.Flags().BoolVarP(&respondOptions.IncludePath, "path", "p", false, "Include the requested path in the output.")
	respondCmd.Flags().BoolVarP(&respondOptions.IncludeHeaders, "headers", "H", false, "Include HTTP request headers in the output.")

	Root.AddCommand(respondCmd)
}
