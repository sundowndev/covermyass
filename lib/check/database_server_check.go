//go:build !windows

package check

type databaseServerCheck struct{}

func NewDatabaseServerCheck() Check {
	return &databaseServerCheck{}
}

func (s *databaseServerCheck) Name() string {
	return "database-server"
}

func (s *databaseServerCheck) Paths() []string {
	return []string{
		"/var/log/mysqld.log",
		"/var/log/mysql.log",
	}
}

func init() {
	AddCheck(NewDatabaseServerCheck())
}
