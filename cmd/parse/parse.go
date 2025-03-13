package parse

import (
	"fmt"
	"os"

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
var outputFile string
var inputFile string
var listFormats bool

func init() {
	Cmd.Flags().StringVarP(&inFormat, "from", "f", "auto", "input format")
	Cmd.Flags().StringVarP(&outFormat, "to", "t", "auto", "output format")
	Cmd.Flags().StringVarP(&query, "query", "q", "", "JMESPath query")
	Cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Filename to write the output to. Omit for stdout.")
	Cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Filename to read input from. Omit for stdin.")
	Cmd.Flags().BoolVarP(&listFormats, "list", "l", false, "List available formats")
}

var Cmd = &cobra.Command{
	Use:     "parse (FILENAME)",
	Short:   "Parse and convert structured data from FILENAME or stdin if omitted",
	Long:    "Parse and convert structured data between different formats. Use --list to see available formats.",
	Aliases: []string{"format"},
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if listFormats {
			fmt.Println("Input formats:")
			for _, name := range parse.Parsers {
				fmt.Printf("  %s\n", name)
			}
			fmt.Println()
			fmt.Println("Output formats:")
			for _, name := range format.Formatters {
				fmt.Printf("  %s\n", name)
			}
			return
		}

		var input []byte
		var err error

		// Handle input from file specified by --input flag or positional argument or stdin
		if inputFile != "" {
			// --input flag takes precedence over positional argument
			input, err = os.ReadFile(inputFile)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("failed to read input file %s: %w", inputFile, err))
			}
		} else {
			// Fall back to positional argument or stdin
			input, err = internal.ReadFileOrStdin(args...)
			if err != nil {
				if len(args) == 0 {
					fmt.Fprintln(os.Stderr, cmd.Short+"\n")
					cmd.Usage()
					os.Exit(1)
				}
				cobra.CheckErr(err)
			}
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

		if outputFile != "" {
			err = os.WriteFile(outputFile, []byte(output), 0644)
			cobra.CheckErr(err)
		} else {
			fmt.Println(output)
		}
	},
}