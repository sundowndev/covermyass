//go:build !windows

package check

type lastLogCheck struct{}

func NewLastLogCheck() Check {
	return &lastLogCheck{}
}

func (s *lastLogCheck) Name() string {
	return "lastlog"
}

func (s *lastLogCheck) Paths() []string {
	return []string{
		"/var/log/lastlog",
	}
}

func init() {
	AddCheck(NewLastLogCheck())
}
