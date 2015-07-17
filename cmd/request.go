package cmd

import (
	"fmt"
	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/util"
	"net/http"
	"os"
)

var RequestAliases map[string]string

func init() {
	RequestAliases = map[string]string{
		"get":    "request",
		"post":   "request",
		"put":    "request",
		"delete": "request",
	}
}

func requestHelp() {
	fmt.Println("Usage: please request <METHOD> [option...] <URL>")
	fmt.Println()
	fmt.Println("Makes a web request to URL using METHOD")
	fmt.Println()
	fmt.Println("Shortcut aliases:")
	for alias := range RequestAliases {
		fmt.Printf("    please %s\n", alias)
	}
	fmt.Println()
	fmt.Println("Input options:")
	fmt.Println("    -i    Include headers from input")
	fmt.Println()
	fmt.Println("Output options:")
	fmt.Println("    -s    Output HTTP status line")
	fmt.Println("    -h    Output headers")
}

func Request(args []string) {
	// Flags
	headers_included := getopt.Bool('i')

	include_headers := getopt.Bool('h')
	include_status := getopt.Bool('s')

	opts := getopt.CommandLine

	opts.SetUsage(requestHelp)

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

	var (
		req *http.Request
		err error
	)

	if termutil.Isatty(os.Stdin.Fd()) {
		req, err = util.CreateRequest(method, url, nil, *headers_included)
	} else {
		req, err = util.CreateRequest(method, url, os.Stdin, *headers_included)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	resp, err := util.GetResponse(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = util.PrintResponse(os.Stdout, resp, *include_headers, *include_status)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
