package formatter

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func Yaml(in interface{}) (out string) {
	bytes, err := yaml.Marshal(in)

	if err != nil {
		fmt.Println("Error generating YAML:", err)
		os.Exit(1)
	}

	return string(bytes)
}

func init() {
	Formatters["yaml"] = Yaml
}
