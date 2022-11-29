//go:build !windows

package check

type mailCheck struct{}

func NewMailCheck() Check {
	return &mailCheck{}
}

func (s *mailCheck) Name() string {
	return "mail"
}

func (s *mailCheck) Paths() []string {
	return []string{
		"/usr/local/psa/var/log/maillog*",
		"/var/log/maillog*",
	}
}

func init() {
	AddCheck(NewMailCheck())
}
