package parse

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/cmd/identify"
	"github.com/stilvoid/please/format"
	"github.com/stilvoid/please/internal"
	"github.com/stilvoid/please/parse"
)

var inFormat string
var outFormat string
var query string

func init() {
	Cmd.Flags().StringVarP(&inFormat, "from", "f", "auto", "input format (see please help parse for formats)")
	Cmd.Flags().StringVarP(&outFormat, "to", "t", "auto", "output format (see please help parse for formats)")
	Cmd.Flags().StringVarP(&query, "query", "q", "", "JMESPath query")

	formats := strings.Builder{}
	formats.WriteString("Input formats:\n")
	for _, name := range parse.Parsers {
		formats.WriteString(fmt.Sprintf("  %s\n", name))
	}
	formats.WriteString("\n")
	formats.WriteString("Output formats:\n")
	for _, name := range format.Formatters {
		formats.WriteString(fmt.Sprintf("  %s\n", name))
	}

	Cmd.Long = "Parse and converted structured data from FILENAME or stdin if omitted.\n\n" + formats.String()
}

var Cmd = &cobra.Command{
	Use:     "parse (FILENAME)",
	Short:   "Parse and convert structured data from FILENAME or stdin if omitted",
	Aliases: []string{"format"},
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		input, err := internal.ReadFileOrStdin(args...)
		if err != nil {
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, cmd.Short+"\n")
				cmd.Usage()
				os.Exit(1)
			}
			cobra.CheckErr(err)
		}

		// Try parsing
		var parsed any

		// Deal with format detection
		if inFormat == "auto" {
			inFormat, parsed, err = identify.Identify(input)
		} else {
			// Try parsing
			parsed, err = parse.Parse(inFormat, input)
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
		output, err := format.Format(outFormat, parsed)
		cobra.CheckErr(err)

		fmt.Println(output)
	},
}