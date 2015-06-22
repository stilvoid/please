package formatter

import (
	"fmt"
	"github.com/clbanning/anyxml"
	"os"
)

func Xml(in interface{}) (out string) {
	bytes, err := anyxml.XmlIndent(in, "", "  ")

	if err != nil {
		fmt.Println("Error generating XML:", err)
		os.Exit(1)
	}

	return string(bytes)
}
