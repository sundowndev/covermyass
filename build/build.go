package build

import (
	"fmt"
	"runtime"
)

var version = "dev"
var commit = "dev"

func Name() string {
	return fmt.Sprintf("%s-%s", version, commit)
}

func String() string {
	return fmt.Sprintf("%s (%s)", Name(), runtime.Version())
}

func IsRelease() bool {
	return Name() != "dev-dev"
}
