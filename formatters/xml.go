package formatters

import (
	"reflect"

	"github.com/clbanning/anyxml"
	"github.com/stilvoid/please/common"
)

func formatXML(in interface{}) (string, error) {
	in = common.ForceStringKeys(in)

	if reflect.TypeOf(in) == nil {
		in = ""
	}

	bytes, err := anyxml.XmlIndent(in, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
