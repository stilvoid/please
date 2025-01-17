package parse

import "fmt"

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
		parser, err := Get(name)

		if err != nil {
			continue
		}

		output, err := parser(input)

		if err != nil {
			continue
		}

		return name, output, nil
	}

	return "", nil, fmt.Errorf("input format could not be identified")
}
