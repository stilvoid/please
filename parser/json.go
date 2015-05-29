package parser

import (
    "encoding/json"
)

func Json(input []byte, path string) (interface{}, error) {
    var parsed interface{}

    err := json.Unmarshal(input, &parsed)

    return parsed, err
}
