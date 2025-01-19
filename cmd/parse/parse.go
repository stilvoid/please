package parse

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"

	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please"
	"github.com/stilvoid/please/cmd/identify"
)

var parserNames []string
var formatNames []string
var inFormat string
var outFormat string
var query string
var listMode bool

func init() {
	parserNames = slices.Collect(maps.Keys(please.Parsers))
	sort.Strings(parserNames)

	formatNames = slices.Collect(maps.Keys(please.Parsers))
	sort.Strings(formatNames)

	Cmd.Flags().StringVarP(&inFormat, "from", "f", "auto", "input format")
	Cmd.Flags().StringVarP(&outFormat, "to", "t", "auto", "output format")
	Cmd.Flags().StringVarP(&query, "query", "q", "", "JMESPath query")
	Cmd.Flags().BoolVar(&listMode, "list", false, "List available input and output types")
}

var Cmd = &cobra.Command{
	Use:   "parse [filename]",
	Short: "Parse and convert structured data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if listMode {
			fmt.Println("Input formats:")
			for _, name := range parserNames {
				fmt.Printf("  %s\n", name)
			}
			fmt.Println()
			fmt.Println("Output formats:")
			for _, name := range formatNames {
				fmt.Printf("  %s\n", name)
			}
			os.Exit(0)
		}

		// Read from stdin
		input, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		// Try parsing
		var parsed any

		// Deal with format detection
		if inFormat == "auto" {
			inFormat, parsed, err = identify.Identify(input)
		} else {
			// Try parsing
			parser, ok := please.Parsers[inFormat]
			if !ok {
				fmt.Fprintf(os.Stderr, "No such parser: %s\n", inFormat)
				os.Exit(1)
			}

			parsed, err = parser(input)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		// Path
		if query != "" {
			parsed, err = jmespath.Search(query, parsed)

			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

		if outFormat == "auto" {
			outFormat = inFormat
		}

		// ...and format back out :)
		formatter, ok := please.Formatters[outFormat]
		if !ok {
			fmt.Fprintf(os.Stderr, "No such formatter: %s\n", outFormat)
			os.Exit(1)
		}

		output, err := formatter(parsed)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(output)
	},
}
