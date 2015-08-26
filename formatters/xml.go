package formatters

import (
	"github.com/clbanning/anyxml"
	"github.com/stilvoid/please/util"
)

func formatXML(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

	bytes, err := anyxml.XmlIndent(in, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
