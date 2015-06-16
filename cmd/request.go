package cmd

import (
	"bufio"
	"bytes"
	"code.google.com/p/getopt"
	"fmt"
	"github.com/andrew-d/go-termutil"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strings"
)

func createRequest(method string, url string, input io.Reader, headers_included bool) *http.Request {
	var req *http.Request
	var headers map[string][]string
	var err error

	method = strings.ToUpper(method)

	if input == nil {
		req, err = http.NewRequest(method, url, nil)
	} else if !headers_included {
		req, err = http.NewRequest(method, url, input)
	} else {
		input_reader := bufio.NewReader(input)

		reader := textproto.NewReader(input_reader)

		headers, err = reader.ReadMIMEHeader()

		if err != nil {
			fmt.Println("Error parsing headers:", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(input_reader)

		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	if input != nil && headers_included {
		for name, values := range headers {
			for i, value := range values {
				if i == 0 {
					req.Header.Set(name, value)
				} else {
					req.Header.Add(name, value)
				}
			}
		}
	}

	return req
}

func getResponse(req *http.Request) *http.Response {
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error contacting host:", err)
		os.Exit(1)
	}

	return resp
}

func printResponse(resp *http.Response, include_headers bool, include_status bool) {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if include_status {
		fmt.Println(resp.StatusCode)
	}

	if include_headers {
		resp.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))
}

func Request(args []string) {
	// Flags
	headers_included := getopt.Bool('i', "Include headers from input")

	include_headers := getopt.Bool('h', "Output headers with the response")
	include_status := getopt.Bool('s', "Output HTTP status line with the response")

	// Cheat because it's better than writing *another* arg parser
	getopt.SetParameters("<url>")
	getopt.SetProgram(fmt.Sprintf("%s <method>", path.Base(os.Args[0])))

	opts := getopt.CommandLine

	// Get the command
	opts.Parse(args)
	if opts.NArgs() < 1 {
		getopt.Usage()
		os.Exit(1)
	}
	method := opts.Arg(0)

	// Deal with flags and get the url
	opts.Parse(opts.Args())
	if opts.NArgs() < 1 {
		getopt.Usage()
		os.Exit(1)
	}
	url := opts.Arg(0)

	var req *http.Request
	if termutil.Isatty(os.Stdin.Fd()) {
		req = createRequest(method, url, nil, *headers_included)
	} else {
		req = createRequest(method, url, os.Stdin, *headers_included)
	}

	resp := getResponse(req)

	printResponse(resp, *include_headers, *include_status)
}
