package cron

// Lock 锁接口：提供Lock() 与 Unlock()方法
type Lock interface {
	Lock()
	UnLock()
}
