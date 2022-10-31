//go:build !windows

package services

import (
	"github.com/sundowndev/covermyass/v2/lib/find"
	"os"
)

type LastLogService struct{}

func NewLastLogService() Service {
	return &LastLogService{}
}

func (s *LastLogService) Name() string {
	return "lastlog"
}

func (s *LastLogService) Paths() []string {
	return []string{
		"/var/log/lastlog",
	}
}

func (s *LastLogService) HandleFile(file find.FileInfo) error {
	return os.Truncate(file.Path(), 0)
}
