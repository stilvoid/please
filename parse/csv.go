package parse

import (
	"bytes"
	"encoding/csv"
)

func Csv(input []byte) (any, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}
