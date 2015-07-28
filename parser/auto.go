package parser

import "fmt"

func Auto(input []byte) (interface{}, error) {
	for _, name := range parseOrder() {
		if parsed, err := Parsers[name].parse(input); err == nil {
			return parsed, err
		}
	}

	return nil, fmt.Errorf("Input format could not be identified")
}

func parseOrder() []string {
	order := make([]string, 0, len(Parsers))

	tried := make(map[string]bool)

	var tryParser func(string)

	tryParser = func(name string) {
		if tried[name] {
			return
		}

		for _, pref := range Parsers[name].prefers {
			tryParser(pref)
		}

		order = append(order, name)
		tried[name] = true
	}

	for name := range Parsers {
		if name != "auto" {
			tryParser(name)
		}
	}

	return order
}

func init() {
	Parsers["auto"] = parser{
		parse: Auto,
	}
}
