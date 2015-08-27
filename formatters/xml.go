package formatters

import (
	"reflect"

	"github.com/clbanning/anyxml"
	"github.com/stilvoid/please/util"
)

func formatXML(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

	if reflect.TypeOf(in) == nil {
		in = ""
	}

	bytes, err := anyxml.XmlIndent(in, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
