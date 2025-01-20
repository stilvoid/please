package respond

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/textproto"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/internal"
)

var headersIncluded bool
var includeHeaders bool
var includeMethod bool
var includeUrl bool
var bodyFn string
var address string
var port int
var status int

func init() {
	Cmd.Flags().BoolVarP(&headersIncluded, "include-headers", "i", false, "Read headers from the response body")
	Cmd.Flags().BoolVarP(&includeHeaders, "output-headers", "o", false, "Output request headers")
	Cmd.Flags().BoolVarP(&includeMethod, "output-method", "m", false, "Output request method")
	Cmd.Flags().BoolVarP(&includeUrl, "output-url", "u", false, "Output request URL")
	Cmd.Flags().StringVarP(&bodyFn, "body", "b", "", "Filename to read the response body from. Use - for stdin.")
	Cmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen on")
	Cmd.Flags().IntVarP(&port, "port", "p", 8000, "Port to listen on")
	Cmd.Flags().IntVarP(&status, "status", "s", 200, "Status code to respond with")
}

var Cmd = &cobra.Command{
	Use:   "respond",
	Short: "Listen for HTTP requests and respond to them",
	Run: func(cmd *cobra.Command, args []string) {
		if headersIncluded && bodyFn == "" {
			cobra.CheckErr(errors.New("You must specify a body filename if --include-headers is set"))
		}

		address = fmt.Sprintf("%s:%d", address, port)

		handler := responder{
			status:          status,
			includeHeaders:  includeHeaders,
			includeMethod:   includeMethod,
			includeUrl:      includeUrl,
			headersIncluded: headersIncluded,
		}

		listener, err := net.Listen("tcp", address)
		cobra.CheckErr(err)

		if bodyFn != "" {
			if bodyFn == "" {
				if termutil.Isatty(os.Stdin.Fd()) {
					cobra.CheckErr(errors.New("Unable to read from stdin"))
				} else {
					handler.data = os.Stdin
				}
			} else {
				var err error
				handler.data, err = os.Open(bodyFn)
				cobra.CheckErr(err)
			}
		}

		server := &http.Server{Addr: address, Handler: handler}

		ch = make(chan bool, 1)

		go func() {
			err := server.Serve(listener)
			if err != nil && err.Error() != "http: Server closed" {
				cobra.CheckErr(err)
			}
		}()

		<-ch

		server.Shutdown(context.Background())
	},
}

var ch chan bool

type responder struct {
	status          int
	includeHeaders  bool
	includeMethod   bool
	includeUrl      bool
	headersIncluded bool
	data            io.ReadSeeker
}

func (h responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	cobra.CheckErr(internal.PrintRequest(req, h.includeMethod, h.includeUrl, h.includeHeaders))

	inputReader := bufio.NewReader(h.data)

	if h.headersIncluded {
		if h.data == nil {
			cobra.CheckErr(errors.New("No body to read headers from"))
		}

		// Parse headers from input
		reader := textproto.NewReader(inputReader)
		headers, err := reader.ReadMIMEHeader()
		cobra.CheckErr(err)

		for name, values := range headers {
			for i, value := range values {
				if i == 0 {
					w.Header().Set(name, value)
				} else {
					w.Header().Add(name, value)
				}
			}
		}
	}

	w.WriteHeader(h.status)

	if h.data != nil {
		io.Copy(w, h.data)
	}

	go func() {
		ch <- true
	}()
}
