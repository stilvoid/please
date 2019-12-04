package version

import (
	"fmt"
	"runtime"
)

var NAME = "Please"
var VERSION = "v1.1.0"

func String() string {
	return fmt.Sprintf("%s %s %s/%s", NAME, VERSION, runtime.GOOS, runtime.GOARCH)
}
