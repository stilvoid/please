package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/pborman/getopt"
	"github.com/stilvoid/please"
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

	format, _, err := identify(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(format)
}

// These should be in order of least to most likely
// i.e. more picky formats should be listed first
var order = []string{
	"xml",
	"mime",
	"json",
	"yaml",
}

// identify tries to figure out the format of the structured data passed in
// If successful, the name of the detected format and a copy of its data parsed into an any will be returned
// If the data format could not be identified, an error will be returned
func identify(input []byte) (string, any, error) {
	for _, name := range order {
		parser, ok := please.Parsers[name]
		if !ok {
			panic(fmt.Errorf("Implementation error. Unknown parser %s", name))
			continue
		}

		output, err := parser(input)
		if err != nil {
			continue
		}

		return name, output, nil
	}

	return "", nil, fmt.Errorf("input format could not be identified")
}
