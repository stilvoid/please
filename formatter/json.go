package formatter

import (
	"fmt"
	"github.com/nytlabs/mxj"
	"os"
)

func Json(in interface{}) (out string) {
	in_map, ok := in.(map[string]interface{})

	if !ok {
		return fmt.Sprintf("\"%s\"", in)
	}

	m := mxj.Map(in_map)

	bytes, err := m.JsonIndent("", "  ")

	if err != nil {
		fmt.Println("Error generating JSON:", err)
		os.Exit(1)
	}

	return string(bytes)
}
