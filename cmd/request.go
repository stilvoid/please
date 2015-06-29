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

	var req *http.Request
	if termutil.Isatty(os.Stdin.Fd()) {
		req = util.CreateRequest(method, url, nil, *headers_included)
	} else {
		req = util.CreateRequest(method, url, os.Stdin, *headers_included)
	}

	resp := util.GetResponse(req)

	util.PrintResponse(resp, *include_headers, *include_status)
}
