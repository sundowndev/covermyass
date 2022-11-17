//go:build !windows

package check

import (
	"github.com/sundowndev/covermyass/v2/lib/find"
	"os"
)

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

func (s *sshdCheck) HandleFile(file find.FileInfo) error {
	return os.Truncate(file.Path(), 0)
}
