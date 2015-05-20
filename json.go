package please

import (
    "encoding/json"
)

func ParseJSON(input []byte, path string) (interface{}, error) {
    var parsed interface{}

    err := json.Unmarshal(input, &parsed)

    return parsed, err
}
