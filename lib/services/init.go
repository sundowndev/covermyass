package services

func Init() {
	AddService(NewSSHdService())
	AddService(NewLastLogService())
	AddService(NewShellHistoryService())
}
