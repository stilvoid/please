package format

import (
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/stilvoid/please/internal"
)

func Toml(in any) (string, error) {
	in = internal.Coerce(in, internal.Config{
		StripNulls: true,
		StringKeys: true,
	})

	b, err := toml.Marshal(in)

	return strings.TrimSpace(string(b)), err
}
