package services

import (
	"github.com/sundowndev/covermyass/v2/lib/find"
)

const (
	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"
)

var services []Service

type Service interface {
	Name() string
	Paths() []string
	HandleFile(find.FileInfo) error
}

func Services() []Service {
	return services
}

func AddService(s Service) {
	services = append(services, s)
}
