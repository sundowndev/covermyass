//go:build !windows

package check

import (
	"github.com/bmatcuk/doublestar/v4"
	"os"
	"testing"
)

func MustGetUserHomeDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}

var shouldMatch = []string{
	"/var/log/mysqld.log",
	"/var/log/mysql.log",
	"/usr/local/psa/var/log/xferlog.any",
	"/usr/local/psa/var/log/xferlog.any",
	"/var/log/xferlog.any",
	"/var/log/secure.any",
	"/var/log/pureftp.log.any",
	"/var/log/apache2/access.log.any",
	"/var/log/apache2/error_log.any",
	"/var/log/httpd.any",
	"/var/log/apache/access.log.any",
	"/var/log/apache/error.log.any",
	"/var/log/nginx/any.log.any",
	"/usr/local/psa/var/log/maillog.log",
	"/var/log/maillog.log",
	MustGetUserHomeDir() + "/.bash_history",
	MustGetUserHomeDir() + "/.zsh_history",
	MustGetUserHomeDir() + "/.node_repl_history",
	MustGetUserHomeDir() + "/.python_history",
	"/var/log/sshd.log",
	"/var/log/lastlog",
	"/var/log/boot.log.any",
	"/var/log/auth.log.any",
	"/var/log/daemon.log.any",
	"/var/log/kern.log.any",
	"/var/log/boot.log.any",
	"/var/log/syslog.any",
	"/var/log/mail.log.any",
	"/var/log/messages.any",
	"/var/log/secure.any",
	"/var/log/btmp.any",
	"/var/log/utmp.any",
	"/var/log/wtmp.any",
	"/var/log/faillog",
	"/var/log/audit/any.log.any",
	"/var/log/dmesg",
}

var shouldNotMatch = []string{
	"/var/log/foo.log",
}

func TestChecks_Valid(t *testing.T) {
	var patterns []string
	for _, c := range checks {
		patterns = append(patterns, c.Paths()...)
	}

	for _, file := range shouldMatch {
		t.Run(file, func(t *testing.T) {
			var match bool
			for _, p := range patterns {
				m, _ := doublestar.Match(p, file)
				if m {
					match = true
					break
				}
			}
			if !match {
				t.Errorf("file %s did not match any pattern", file)
			}
		})
	}
}

func TestChecks_Invalid(t *testing.T) {
	var patterns []string
	for _, c := range checks {
		patterns = append(patterns, c.Paths()...)
	}

	for _, file := range shouldNotMatch {
		t.Run(file, func(t *testing.T) {
			var match bool
			for _, p := range patterns {
				m, _ := doublestar.Match(p, file)
				if m {
					match = true
					break
				}
			}
			if match {
				t.Errorf("file %s unexpectedly matched one or more patterns", file)
			}
		})
	}
}
