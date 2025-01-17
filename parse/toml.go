package parse

import "github.com/BurntSushi/toml"

func Toml(input []byte) (any, error) {
	var parsed any

	err := toml.Unmarshal(input, &parsed)

	return parsed, err
}
