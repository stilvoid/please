package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/textproto"
	"os"
	"strconv"

	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/common"
)

type responder struct {
	status          int
	includeHeaders  bool
	includeMethod   bool
	includeUrl      bool
	headersIncluded bool
	listener        net.Listener
	data            io.ReadSeeker
}

func init() {
	Commands["respond"] = respondCommand
}

func respondHelp() {
	fmt.Println("Usage: please respond [option...] <STATUS> [<ADDRESS>[:<PORT>]]")
	fmt.Println()
	fmt.Println("Listens on the specified address and port and responds with the chosen status code.")
	fmt.Println("Any data on stdin will be used as the body of the response.")
	fmt.Println("The request body will be printed to stdout.")
	fmt.Println()
	fmt.Println("Input options:")
	fmt.Println("    -i    Include headers from input")
	fmt.Println()
	fmt.Println("Output options:")
	fmt.Println("    -m    Output the request method")
	fmt.Println("    -u    Output the requested path")
	fmt.Println("    -h    Output headers with the request")
}

func (h responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer h.listener.Close()

	err := common.WriteRequest(os.Stdout, req, h.includeMethod, h.includeUrl, h.includeHeaders)
	os.Stdout.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	inputReader := bufio.NewReader(h.data)

	if h.headersIncluded {
		if h.data == nil {
			fmt.Println("Error reading headers")
			os.Exit(1)
		}

		// Parse headers from input
		reader := textproto.NewReader(inputReader)
		headers, err := reader.ReadMIMEHeader()

		if err != nil {
			fmt.Println("Error parsing headers:", err)
			os.Exit(1)
		}

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
		io.Copy(w, inputReader)
	}
}

func respondCommand(args []string) {
	// Flags
	headersIncluded := getopt.Bool('i')

	includeHeaders := getopt.Bool('h')
	includeMethod := getopt.Bool('m')
	includeUrl := getopt.Bool('u')

	opts := getopt.CommandLine

	opts.SetUsage(respondHelp)

	// Deal with flags and get the url
	opts.Parse(args)
	if opts.NArgs() < 1 {
		getopt.Usage()
		os.Exit(1)
	}

	status, err := strconv.Atoi(opts.Arg(0))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid status code: %s\n", opts.Arg(0))
		os.Exit(1)
	}

	address := "0.0.0.0:8000"

	if opts.NArgs() >= 2 {
		address = opts.Arg(1)
	}

	handler := responder{
		status:          status,
		includeHeaders:  *includeHeaders,
		includeMethod:   *includeMethod,
		includeUrl:      *includeUrl,
		headersIncluded: *headersIncluded,
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	handler.listener = listener

	if termutil.Isatty(os.Stdin.Fd()) {
		handler.data = nil
	} else {
		handler.data = os.Stdin
	}

	server := &http.Server{Addr: address, Handler: handler}

	server.Serve(listener)

}
