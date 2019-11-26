package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/andrew-d/go-termutil"
	"github.com/spf13/cobra"
	"github.com/stilvoid/please/parsers"
)

var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "Identify a structured data format",
	Long: fmt.Sprintf(`Identifies the format of the structured data on stdin

Supported formats:
  %s`, strings.Join(parsers.Names(), "\n  ")),
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if termutil.Isatty(os.Stdin.Fd()) {
			panic("No data on stdin.")
		}

		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(fmt.Errorf("Error reading input: %s", err.Error()))
		}

		format, _, err := parsers.Identify(input)
		if err != nil {
			panic(err)
		}

		fmt.Println(format)
	},
}

func init() {
	Root.AddCommand(identifyCmd)
}
