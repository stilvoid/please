package cmd

import (
	"fmt"
	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/util"
	"net/http"
	"os"
	"path"
)

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
