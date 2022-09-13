package cron

type Lock interface {
	Lock()
	UnLock()
}
