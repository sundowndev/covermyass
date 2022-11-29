package check

var checks []Check

type Check interface {
	Name() string
	Paths() []string
}

func GetAllChecks() []Check {
	return checks
}

func AddCheck(s Check) {
	checks = append(checks, s)
}
