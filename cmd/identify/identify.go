package identify

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stilvoid/please/internal"
	"github.com/stilvoid/please/parse"
)

var Cmd = &cobra.Command{
	Use:   "identify (FILENAME)",
	Short: "Identify the format of some structured data from FILENAME or stdin if omitted",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := internal.ReadFileOrStdin(args...)
		if err != nil {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
			cobra.CheckErr(err)
		}

		format, _, err := Identify(data)
		cobra.CheckErr(err)

		fmt.Println(format)
	},
}

// These should be in order of least to most likely
// i.e. more picky formats should be listed first
var order = []string{
	"xml",
	"mime",
	"json",
	"toml",
	"yaml",
}

// Identify tries to figure out the format of the structured data passed in
// If successful, the name of the detected format and a copy of its data parsed into an any will be returned
// If the data format could not be identified, an error will be returned
func Identify(input []byte) (string, any, error) {
	for _, name := range order {
		output, err := parse.Parse(name, input)
		if err != nil {
			continue
		}

		return name, output, nil
	}

	return "", nil, fmt.Errorf("input format could not be identified")
}
