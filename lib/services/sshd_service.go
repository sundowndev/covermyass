//go:build !windows

package services

import (
	"github.com/sundowndev/covermyass/v2/lib/find"
	"os"
)

type SSHdService struct{}

func NewSSHdService() Service {
	return &SSHdService{}
}

func (s *SSHdService) Name() string {
	return "sshd"
}

func (s *SSHdService) Paths() []string {
	return []string{
		"/var/log/sshd.log",
	}
}

func (s *SSHdService) HandleFile(file find.FileInfo) error {
	return os.Truncate(file.Path(), 0)
}
