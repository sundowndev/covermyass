package services

func Init() {
	AddService(NewSSHdService())
}
