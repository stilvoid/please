package please

import (
    "bytes"
    "encoding/csv"
)

func ParseCSV(input []byte, path string) (interface{}, error) {
    return csv.NewReader(bytes.NewReader(input)).ReadAll()
}
