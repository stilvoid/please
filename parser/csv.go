package parser

import (
	"bytes"
	"encoding/csv"
)

func parseCSV(input []byte) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}

func init() {
	parsers["csv"] = parser{
		parse:   parseCSV,
		prefers: []string{"yaml"},
	}
}
