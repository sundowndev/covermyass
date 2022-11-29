//go:build !windows

package check

type ftpCheck struct{}

func NewFTPCheck() Check {
	return &ftpCheck{}
}

func (s *ftpCheck) Name() string {
	return "ftp"
}

func (s *ftpCheck) Paths() []string {
	return []string{
		"/usr/local/psa/var/log/xferlog*",
		"/var/log/xferlog*",
		"/var/log/secure*",
		"/var/log/pureftp.log*",
	}
}

func init() {
	AddCheck(NewFTPCheck())
}
