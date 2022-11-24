package parse

import (
	"bytes"
	"encoding/csv"
)

func Csv(input []byte) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}
