package check

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/lib/find"
	"os"
)

type shellHistoryCheck struct{}

func NewShellHistoryCheck() Check {
	return &shellHistoryCheck{}
}

func (s *shellHistoryCheck) Name() string {
	return "shell_history"
}

func (s *shellHistoryCheck) Paths() []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.Error(err)
		return []string{}
	}
	return []string{
		fmt.Sprintf("%s/.bash_history", homeDir),
		fmt.Sprintf("%s/.zsh_history", homeDir),
		fmt.Sprintf("%s/.node_repl_history", homeDir),
		fmt.Sprintf("%s/.python_history", homeDir),
	}
}

func (s *shellHistoryCheck) HandleFile(file find.FileInfo) error {
	return os.Truncate(file.Path(), 0)
}

func init() {
	AddCheck(NewShellHistoryCheck())
}
