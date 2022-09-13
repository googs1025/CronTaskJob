package cron

import (
	"crontask/tasks"
	"log"
	"os"
)

type TaskScheduler struct {
	tasks  []tasks.TaskInterface
	swap   []tasks.TaskInterface
	add    chan tasks.TaskInterface
	remove chan string
	stop   chan struct{}
	Logger tasks.TaskLogInterface
	lock   bool

}

type OnceCron struct {
	*TaskScheduler
}

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

func NewCron() *OnceCron {
	return &OnceCron{
		TaskScheduler: NewScheduler(),
	}
}
