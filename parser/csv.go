package parser

import (
	"bytes"
	"encoding/csv"
)

func parseCsv(input []byte) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}

func init() {
	parsers["csv"] = parser{
		parse:   parseCsv,
		prefers: []string{"yaml"},
	}
}
