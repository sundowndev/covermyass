package check

func Init() {
	AddCheck(NewSSHdCheck())
	AddCheck(NewLastLogCheck())
	AddCheck(NewShellHistoryCheck())
}
