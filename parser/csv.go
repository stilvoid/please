package parser

import (
	"bytes"
	"encoding/csv"
)

func Csv(input []byte) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}

func init() {
	Parsers["csv"] = parser{
		parse:   Csv,
		prefers: []string{"yaml"},
	}
}
