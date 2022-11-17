//go:build !windows

package check

import (
	"github.com/sundowndev/covermyass/v2/lib/find"
	"os"
)

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

func (s *lastLogCheck) HandleFile(file find.FileInfo) error {
	return os.Truncate(file.Path(), 0)
}
