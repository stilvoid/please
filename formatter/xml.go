package formatter

import (
	"github.com/clbanning/anyxml"
)

func formatXML(in interface{}) (string, error) {
	bytes, err := anyxml.XmlIndent(in, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func init() {
	formatters["xml"] = formatXML
}
