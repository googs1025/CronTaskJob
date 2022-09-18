package cron

import (
	"crontask/tasks"
	"log"
	"os"
)

// TaskScheduler 是调度对列对象。
type TaskScheduler struct {
	// 主要放任务对象的list
	tasks  []tasks.TaskInterface
	swap   []tasks.TaskInterface

	add    chan tasks.TaskInterface
	//
	remove chan string
	stop   chan struct{}
	// 提供日志
	Logger tasks.TaskLogInterface
	lock   bool

}

// OnceCron 一次调度的对象
type OnceCron struct {
	*TaskScheduler
}

// NewScheduler 调度对列的构建函数
func NewScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]tasks.TaskInterface, 0),
		swap: make([]tasks.TaskInterface, 0),
		add: make(chan tasks.TaskInterface),
		stop: make(chan struct{}),
		remove: make(chan string),
		Logger: log.New(os.Stdout, "[Control]:", log.Ldate|log.Ltime|log.Lshortfile),
	}

}

// NewCron 一次性调度对列的构建函数
func NewCron() *OnceCron {
	return &OnceCron{
		TaskScheduler: NewScheduler(),
	}
}
