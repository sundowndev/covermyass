//go:build !windows

package check

type httpServerCheck struct{}

func NewHTTPServerCheck() Check {
	return &httpServerCheck{}
}

func (s *httpServerCheck) Name() string {
	return "http-server"
}

func (s *httpServerCheck) Paths() []string {
	return []string{
		"/var/log/apache2/access.log*",
		"/var/log/apache2/error_log*",
		"/var/log/httpd*",
		"/var/log/apache/access.log*",
		"/var/log/apache/error.log*",
		"/var/log/nginx/*.log*",
	}
}

func init() {
	AddCheck(NewHTTPServerCheck())
}
