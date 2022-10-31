package services

import "os"

const (
	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"
)

var services []Service

type Service interface {
	Name() string
	Paths() map[string][]string
	HandleFile(string, os.FileInfo) error
}

func Services() []Service {
	return services
}

func AddService(s Service) {
	services = append(services, s)
}
