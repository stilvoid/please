package cmd

import (
	"bufio"
	"code.google.com/p/getopt"
	"fmt"
	"github.com/andrew-d/go-termutil"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/textproto"
	"os"
)

type responder struct {
	status           string
	include_headers  bool
	include_method   bool
	include_url      bool
	headers_included bool
	listener         net.Listener
	data             io.ReadSeeker
}

func printRequest(req *http.Request, include_method bool, include_url bool, include_headers bool) {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if include_method {
		fmt.Println(req.Method)
	}

	if include_url {
		fmt.Println(req.URL)
	}

	if include_headers {
		req.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))
}

func (h responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	printRequest(req, h.include_method, h.include_url, h.include_headers)

	input_reader := bufio.NewReader(h.data)

	if h.headers_included {
		if h.data == nil {
			fmt.Println("Error reading headers")
			os.Exit(1)
		}

		// Parse headers from input
		reader := textproto.NewReader(input_reader)
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

	if h.data != nil {
		io.Copy(w, input_reader)
	}

	h.listener.Close()
}

func Respond(args []string) {
	// Flags
	headers_included := getopt.Bool('i', "Output headers with the response")

	include_headers := getopt.Bool('h', "Include headers in output")
	include_method := getopt.Bool('m', "Include method in output")
	include_url := getopt.Bool('u', "Include URL in output")

	// Cheat because it's better than writing *another* arg parser
	getopt.SetParameters("<status> [<address>[:<port>]]")

	opts := getopt.CommandLine

	// Deal with flags and get the url
	opts.Parse(args)
	if opts.NArgs() < 1 {
		getopt.Usage()
		os.Exit(1)
	}

	status := opts.Arg(0)

	address := "0.0.0.0:8000"

	if opts.NArgs() >= 2 {
		address = opts.Arg(1)
	}

	handler := responder{
		status:           status,
		include_headers:  *include_headers,
		include_method:   *include_method,
		include_url:      *include_url,
		headers_included: *headers_included,
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
