package formatter

import "github.com/stilvoid/please/util"

type Formatter func(interface{}) string

var Formatters = map[string]Formatter{
	"bash": Bash,
	"dot":  Dot,
	"json": Json,
	"xml":  Xml,
	"yaml": Yaml,
}

func Format(input interface{}, format string) string {
	if format != "yaml" {
		// Pretty much everything hates non-string keys :S
		input = util.ForceStringKeys(input)
	}

	return Formatters[format](input)
}
