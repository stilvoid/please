package parser

import (
	"fmt"
)

type Parser func([]byte) (interface{}, error)

var Parsers = map[string]Parser{
	"csv":  Csv,
	"html": Html,
	"json": Json,
	"mime": Mime,
	"xml":  Xml,
	"yaml": Yaml,
}

var preference = []string{
	"json",
	"xml",
	"yaml",
	"csv",
	"html",
	"mime",
}

func Parse(input []byte, format string) (interface{}, string, error) {
	if format == "auto" {
		for _, name := range preference {
			parsed, err := Parsers[name](input)

			if err == nil {
				return parsed, name, err
			}
		}

		return nil, "", fmt.Errorf("Input format could not be identified")
	}

	parser, ok := Parsers[format]

	if !ok {
		return nil, "", fmt.Errorf("No such parser: %s", format)
	}

	output, err := parser(input)

	return output, format, err
}
