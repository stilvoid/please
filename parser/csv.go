package parser

import (
	"bytes"
	"encoding/csv"
)

func Csv(input []byte, path string) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}
