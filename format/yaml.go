package format

import "gopkg.in/yaml.v2"

func formatYAML(in interface{}) (string, error) {
	bytes, err := yaml.Marshal(in)

	// We strip off the trailing newline that yaml.v2 seems to insist on
	return string(bytes[:len(bytes)-1]), err
}
