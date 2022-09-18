package tasks

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Job 接口：实现Run()方法
type Job interface {
	Run()
}

// Run 方法：实现func() 对象
func (f FuncJob) Run() {
	f()
}

// FuncJob 特定函数对象
type FuncJob func()

// Task 对象
type Task struct {
	// 实现Job接口
	Job     Job
	// 任务元数据
	// TODO: 可以把记录放入mysql中。
	ID      string
	RunTime int64
	Spacing int64
	EndTime int64
	Number  int
}

// 这里的取得就是建立并返回Task对象

// GetTaskWithFunc 取得task对象
func GetTaskWithFunc(unixTime int64, f func()) *Task {
	return &Task{
		Job:     FuncJob(f),
		RunTime: unixTime,
		ID:      uuid.New().String(),
	}
}

// GetTaskWithFuncSpacingNumber 取得task对象
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

// GetTaskWithFuncSpacing 取得task对象
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

