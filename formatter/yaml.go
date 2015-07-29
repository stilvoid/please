package formatter

import (
	"gopkg.in/yaml.v2"
)

func formatYAML(in interface{}) (string, error) {
	bytes, err := yaml.Marshal(in)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func init() {
	formatters["yaml"] = formatYAML
}
