package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/andrew-d/go-termutil"
	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/formatters"
	"github.com/stilvoid/please/parsers"
)

var inFormat string
var outFormat string
var filter string

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Pretty-print, convert, and filter data structures",
	Long: fmt.Sprintf(`Pretty-print, convert, and filter data structures.
Data will be read from stdin and output will be written to stdout.

Supported input formats:
  auto
  %s

Supported output formats:
  auto
  %s`, strings.Join(parsers.Names(), "\n  "), strings.Join(formatters.Names(), "\n  ")),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse the filter
		var jp *jmespath.JMESPath
		if filter != "" {
			jp = jmespath.MustCompile(filter)
		}

		if termutil.Isatty(os.Stdin.Fd()) {
			// Don't try to read from the terminal
			panic("No data on stdin.")
		}

		// Read the input
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(fmt.Errorf("Unable to read input: %s", err.Error()))
		}

		// Try parsing
		var parsed interface{}

		// Deal with format detection
		if inFormat == "auto" {
			inFormat, parsed, err = parsers.Identify(input)
		} else {
			// Try parsing
			parser, err := parsers.Get(inFormat)
			if err != nil {
				panic(err)
			}

			parsed, err = parser(input)
		}

		if err != nil {
			panic(err)
		}

		// Path
		if jp != nil {
			parsed, err = jp.Search(parsed)
			if err != nil {
				panic(err)
			}
		}

		if outFormat == "auto" {
			outFormat = inFormat
		}

		// ...and format back out :)
		formatter, err := formatters.Get(outFormat)
		if err != nil {
			panic(err)
		}

		output, err := formatter(parsed)
		if err != nil {
			panic(err)
		}

		fmt.Println(output)
	},
}

func init() {
	parseCmd.Flags().StringVarP(&inFormat, "input", "i", "auto", "Input format.")
	parseCmd.Flags().StringVarP(&outFormat, "output", "o", "auto", "Output format.")
	parseCmd.Flags().StringVarP(&filter, "filter", "f", "", "A JMESPath query to perform before outputting.")
	Root.AddCommand(parseCmd)
}
