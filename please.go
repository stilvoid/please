package please

import (
	"github.com/stilvoid/please/format"
	"github.com/stilvoid/please/parse"
)

// Type Formatter is a function that takes an any and attempts to format it as a string
type Formatter func(any) (string, error)

// Formatters collects all formatters available to please
var Formatters = map[string]Formatter{
	"bash":  format.Bash,
	"dot":   format.Dot,
	"json":  format.Json,
	"xml":   format.Xml,
	"yaml":  format.Yaml,
	"query": format.Query,
	"toml":  format.Toml,
}

// Type Parser is a function that takes a byte slice and attempts to parse it into a structure format in an any
type Parser func([]byte) (any, error)

// Parser collects all parsers available to please
var Parsers = map[string]Parser{
	"csv":   parse.Csv,
	"html":  parse.Html,
	"json":  parse.Json,
	"mime":  parse.Mime,
	"xml":   parse.Xml,
	"yaml":  parse.Yaml,
	"query": parse.Query,
	"toml":  parse.Toml,
}
