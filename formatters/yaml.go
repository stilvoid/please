package formatters

import (
	"github.com/stilvoid/please/util"
	"gopkg.in/yaml.v2"
)

func YAML(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

	bytes, err := yaml.Marshal(in)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
