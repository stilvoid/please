package parser

import "fmt"

// Identify tries to figure out the format of the structured data passed in
// If the data format could not be identified, an error will be returned
func Identify(input []byte) (string, error) {
	for _, name := range parseOrder() {
		if parsed, err := parsers[name].parse(input); err == nil {
			fmt.Println(name, parsed)

			return name, nil
		}
	}

	return "", fmt.Errorf("input format could not be identified")
}

func auto(input []byte) (interface{}, error) {
	for _, name := range parseOrder() {
		if parsed, err := parsers[name].parse(input); err == nil {
			return parsed, err
		}
	}

	return nil, fmt.Errorf("input format could not be identified")
}

func parseOrder() []string {
	order := make([]string, 0, len(parsers))

	tried := make(map[string]bool)

	var tryParser func(string)

	tryParser = func(name string) {
		if tried[name] {
			return
		}

		for _, pref := range parsers[name].prefers {
			tryParser(pref)
		}

		order = append(order, name)
		tried[name] = true
	}

	for name := range parsers {
		if name != "auto" {
			tryParser(name)
		}
	}

	return order
}

func init() {
	parsers["auto"] = parser{
		parse: auto,
	}
}
