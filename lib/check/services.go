package check

import (
	"github.com/sundowndev/covermyass/v2/lib/find"
)

const (
	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"
)

var checks []Check

type Check interface {
	Name() string
	Paths() []string
	HandleFile(find.FileInfo) error
}

func GetAllChecks() []Check {
	return checks
}

func AddCheck(s Check) {
	checks = append(checks, s)
}
