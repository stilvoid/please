package parsers

import (
	"bytes"
	"encoding/csv"
)

func parseCSV(input []byte) (interface{}, error) {
	return csv.NewReader(bytes.NewReader(input)).ReadAll()
}
