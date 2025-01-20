package identify

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stilvoid/please"
)

var Cmd = &cobra.Command{
	Use:   "identify [filename]",
	Short: "Identify the format of some structured data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(args[0])
		cobra.CheckErr(err)

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
	"yaml",
}

// Identify tries to figure out the format of the structured data passed in
// If successful, the name of the detected format and a copy of its data parsed into an any will be returned
// If the data format could not be identified, an error will be returned
func Identify(input []byte) (string, any, error) {
	for _, name := range order {
		output, err := please.Parse(name, input)
		if err != nil {
			continue
		}

		return name, output, nil
	}

	return "", nil, fmt.Errorf("input format could not be identified")
}
