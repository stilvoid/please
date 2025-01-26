package parse

import (
	"fmt"

	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please"
	"github.com/stilvoid/please/cmd/identify"
	"github.com/stilvoid/please/internal"
)

var inFormat string
var outFormat string
var query string
var listMode bool

func init() {
	Cmd.Flags().StringVarP(&inFormat, "from", "f", "auto", "input format")
	Cmd.Flags().StringVarP(&outFormat, "to", "t", "auto", "output format")
	Cmd.Flags().StringVarP(&query, "query", "q", "", "JMESPath query")
	Cmd.Flags().BoolVarP(&listMode, "list", "l", false, "List available input and output types")
}

var Cmd = &cobra.Command{
	Use:   "parse (FILENAME)",
	Short: "Parse and convert structured data from FILENAME or stdin if omitted",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if listMode {
			fmt.Println("Input formats:")
			for _, name := range please.Parsers {
				fmt.Printf("  %s\n", name)
			}
			fmt.Println()
			fmt.Println("Output formats:")
			for _, name := range please.Formatters {
				fmt.Printf("  %s\n", name)
			}
			return
		}

		// Read from stdin?
		input, err := internal.ReadFileOrStdin(args...)
		cobra.CheckErr(err)

		// Try parsing
		var parsed any

		// Deal with format detection
		if inFormat == "auto" {
			inFormat, parsed, err = identify.Identify(input)
		} else {
			// Try parsing
			parsed, err = please.Parse(inFormat, input)
		}

		cobra.CheckErr(err)

		// Path
		if query != "" {
			parsed, err = jmespath.Search(query, parsed)
			cobra.CheckErr(err)
		}

		if outFormat == "auto" {
			outFormat = inFormat
		}

		// ...and format back out :)
		output, err := please.Format(outFormat, parsed)
		cobra.CheckErr(err)

		fmt.Println(output)
	},
}
