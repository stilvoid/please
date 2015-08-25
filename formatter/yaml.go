package formatter

import (
	"github.com/stilvoid/please/util"
	"gopkg.in/yaml.v2"
)

func formatYAML(in interface{}) (string, error) {
	in = util.ForceStringKeys(in)

	bytes, err := yaml.Marshal(in)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func init() {
	formatters["yaml"] = formatYAML
}
