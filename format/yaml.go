package format

import (
	"github.com/stilvoid/please/internal"
	"gopkg.in/yaml.v2"
)

func Yaml(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{})

	bytes, err := yaml.Marshal(in)

	// We strip off the trailing newline that yaml.v2 seems to insist on
	return string(bytes[:len(bytes)-1]), err
}
