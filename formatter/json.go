package formatter

import (
	"fmt"
	"github.com/nytlabs/mxj"
	"os"
)

func formatJson(in interface{}) (out string) {
	inMap, ok := in.(map[string]interface{})

	if !ok {
		return fmt.Sprintf("\"%s\"", in)
	}

	m := mxj.Map(inMap)

	bytes, err := m.JsonIndent("", "  ")

	if err != nil {
		fmt.Println("Error generating JSON:", err)
		os.Exit(1)
	}

	return string(bytes)
}

func init() {
	formatters["json"] = formatJson
}
