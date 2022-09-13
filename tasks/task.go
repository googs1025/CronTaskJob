package tasks

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Job interface {
	Run()
}


func (f FuncJob) Run() {
	f()
}

type FuncJob func()


type Task struct {
	Job     Job
	ID      string
	RunTime int64
	Spacing int64
	EndTime int64
	Number  int
}



func GetTaskWithFunc(unixTime int64, f func()) *Task {
	return &Task{
		Job:     FuncJob(f),
		RunTime: unixTime,
		ID:      uuid.New().String(),
	}
}

func GetTaskWithFuncSpacingNumber(spacing int64, number int, f func()) *Task {
	return &Task{
		Job:     FuncJob(f),
		Spacing: spacing,
		RunTime: time.Now().UnixNano() + spacing,
		Number:  number,
		EndTime: time.Now().UnixNano() + int64(number)*spacing*int64(time.Second),
		ID:      uuid.New().String(),
	}
}

func GetTaskWithFuncSpacing(spacing int64, endTime int64,f func()) *Task {
	return &Task{
		Job:     FuncJob(f),
		Spacing: spacing,
		RunTime: time.Now().UnixNano() + spacing,
		EndTime: endTime,
		ID:      uuid.New().String(),
	}
}

func (task *Task) toString() string {
	return fmt.Sprintf("uuid: %s, runTime: %d, spacing: %d, endTime: %d, number: %d", task.ID, task.RunTime, task.Spacing, task.EndTime, task.Number)
}

