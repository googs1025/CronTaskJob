package tasks

type TaskInterface interface {
	TaskGetInterface
	TaskSetInterface
}

type TaskGetInterface interface {
	RunJob()
	GetJob() Job
	GetUuid() string
	GetRunTime() int64
	GetSpacing() int64
	GetEndTime() int64
	GetNumber() int

}

type TaskSetInterface interface {
	SetJob(job Job) TaskSetInterface
	SetRunTime(runtime int64) TaskSetInterface
	SetID(uuid string) TaskSetInterface
	SetSpacing(spacing int64) TaskSetInterface
	SetEndTime(endtime int64) TaskSetInterface
	SetRunNumber(number int) TaskSetInterface
}

type TaskLogInterface interface {
	Println(v ...interface{})
}

