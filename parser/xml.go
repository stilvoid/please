package parser

import (
    "github.com/clbanning/x2j"
)

func Xml(input []byte, path string) (interface{}, error) {
    parsed := make(map[string]interface{})

    err := x2j.Unmarshal(input, &parsed)

    return parsed, err
}
