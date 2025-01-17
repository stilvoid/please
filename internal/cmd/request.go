package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/internal/web"
)

var requestAliases = []string{
	"get",
	"post",
	"put",
	"delete",
}

func init() {
	Commands["request"] = requestCommand

	for _, alias := range requestAliases {
		Aliases[alias] = "request"
	}
}

func requestHelp() {
	fmt.Println("Usage: please request <METHOD> [option...] <URL>")
	fmt.Println()
	fmt.Println("Makes a web request to URL using METHOD")
	fmt.Println()
	fmt.Println("Shortcut aliases:")
	for _, alias := range requestAliases {
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

func requestCommand(args []string) {
	// Flags
	headersIncluded := getopt.Bool('i')

	includeHeaders := getopt.Bool('h')
	includeStatus := getopt.Bool('s')

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

	var input io.Reader

	if !termutil.Isatty(os.Stdin.Fd()) {
		input = os.Stdin
	}

	resp, err := web.MakeRequest(method, url, input, *headersIncluded)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = web.WriteResponse(os.Stdout, resp, *includeHeaders, *includeStatus)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
