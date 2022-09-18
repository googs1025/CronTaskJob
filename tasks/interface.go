package tasks

// TaskInterface 接口：分别实现TaskGetInterface TaskSetInterface接口
type TaskInterface interface {
	TaskGetInterface
	TaskSetInterface
}

// TaskGetInterface 接口：实现多种Get方法与最重要的RunJob()执行方法
type TaskGetInterface interface {
	RunJob()
	GetJob() Job
	GetUUID() string
	GetRunTime() int64
	GetSpacing() int64
	GetEndTime() int64
	GetNumber() int

}

// TaskSetInterface 接口：实现多种Set方法
type TaskSetInterface interface {
	SetJob(job Job) TaskSetInterface
	SetRunTime(runtime int64) TaskSetInterface
	SetID(uuid string) TaskSetInterface
	SetSpacing(spacing int64) TaskSetInterface
	SetEndTime(endtime int64) TaskSetInterface
	SetRunNumber(number int) TaskSetInterface
}

// TaskLogInterface 接口：实现打印日志方法
type TaskLogInterface interface {
	Println(v ...interface{})
}

