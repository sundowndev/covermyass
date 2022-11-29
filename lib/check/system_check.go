//go:build !windows

package check

type systemCheck struct{}

func NewSystemCheck() Check {
	return &systemCheck{}
}

func (s *systemCheck) Name() string {
	return "system"
}

func (s *systemCheck) Paths() []string {
	return []string{
		"/var/log/lastlog",
		"/var/log/boot.log*",
		"/var/log/auth.log*",
		"/var/log/daemon.log*",
		"/var/log/kern.log*",
		"/var/log/boot.log*",
		"/var/log/syslog*",
		"/var/log/mail.log*",
		"/var/log/messages*",
		"/var/log/secure*",
		"/var/log/btmp*",
		"/var/log/utmp*",
		"/var/log/wtmp*",
		"/var/log/faillog",
		"/var/log/audit/*.log*",
		"/var/log/dmesg",
	}
}

func init() {
	AddCheck(NewSystemCheck())
}
