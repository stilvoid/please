package parsers

import (
	"bytes"
	"encoding/csv"
)

func CSV(input []byte) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}
