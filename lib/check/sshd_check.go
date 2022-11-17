//go:build !windows

package check

type sshdCheck struct{}

func NewSSHdCheck() Check {
	return &sshdCheck{}
}

func (s *sshdCheck) Name() string {
	return "sshd"
}

func (s *sshdCheck) Paths() []string {
	return []string{
		"/var/log/sshd.log",
	}
}

func init() {
	AddCheck(NewSSHdCheck())
}
