package parser

import "fmt"

func auto(input []byte) (interface{}, error) {
	for _, name := range parseOrder() {
		if parsed, err := parsers[name].parse(input); err == nil {
			return parsed, err
		}
	}

	return nil, fmt.Errorf("Input format could not be identified")
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
