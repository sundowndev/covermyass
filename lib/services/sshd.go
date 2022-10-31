package services

import "os"

type SSHdService struct{}

func NewSSHdService() Service {
	return &SSHdService{}
}

func (s *SSHdService) Name() string {
	return "sshd"
}

func (s *SSHdService) Paths() map[string][]string {
	return map[string][]string{
		Linux: {
			"/var/log/sshd.log",
		},
	}
}

func (s *SSHdService) HandleFile(path string, info os.FileInfo) error {
	return os.Truncate(path, 0)
}
