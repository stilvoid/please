package parse

import "github.com/clbanning/mxj/x2j"

func Xml(input []byte) (any, error) {
	return x2j.XmlToMap(input)
}
