package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please/parsers"
)

func init() {
	Commands["identify"] = identifyCommand
}

func identifyHelp() {
	fmt.Println("Usage: please identify")
	fmt.Println()
	fmt.Println("Identifies the format of the structured data on standard input")
}

func identifyCommand(args []string) {
	opts := getopt.CommandLine

	opts.SetUsage(identifyHelp)

	if termutil.Isatty(os.Stdin.Fd()) {
		getopt.Usage()
		os.Exit(1)
	}

	var err error

	// Read from stdin
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(1)
	}

	format, _, err := parsers.Identify(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(format)
}
