package main

import (
    "code.google.com/p/getopt"
    "fmt"
    "io/ioutil"
    "offend.me.uk/please"
    "os"
)

var parsers map[string]func([]byte, string) (interface{}, error)
var formatters map[string]func(interface{}, string) string

func parseAuto(input []byte, path string) (interface{}, error) {
    var parsed interface{}
    var err error

    for name, parser := range(parsers) {
        fmt.Println(name)
        if name != "auto" {
            parsed, err = parser(input, path)

            if err == nil {
                break
            }
        }
    }

    return parsed, err
}

func init() {
    parsers = map[string]func([]byte, string) (interface{}, error) {
        "auto": parseAuto,
        "json": please.ParseJSON,
        "xml": please.ParseXML,
        "csv": please.ParseCSV,
        "html": please.ParseHTML,
    }

    formatters = map[string]func(interface{}, string) string {
        "bash": please.FormatBash,
    }
}

func main() {
    // Flags
    in_format := getopt.String('i', "auto", "Parse the input as 'types'", "type")
    out_format := getopt.String('o', "bash", "Use 'type' as the output format", "type")
    getopt.Parse()

    // Validate parser
    if _, ok := parsers[*in_format]; !ok {
        fmt.Printf("Unknown input format: %s\n", *in_format)
        os.Exit(1)
    }

    // Validate formatter
    if _, ok := formatters[*out_format]; !ok {
        fmt.Printf("Unknown output format: %s\n", *out_format)
        os.Exit(1)
    }

    var err error

    // Read from stdin
    input, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Println("Error reading input")
        os.Exit(1)
    }

    // Path
    var path string

    if getopt.NArgs() > 0 {
        path = getopt.Arg(0)
    }

    // Try parsing
    parsed, err := parsers[*in_format](input, path)

    if err != nil {
        fmt.Fprintln(os.Stderr, "Input could not be parsed")
        fmt.Println(err)
        os.Exit(1)
    }

    // ...and format back out :)
    fmt.Println(formatters[*out_format](parsed, path))
}
