package tasks

func (task *Task) SetJob(job Job) TaskSetInterface {
	task.Job = job
	return task
}

func (task *Task) SetRunTime(runtime int64) TaskSetInterface {

	if runtime < 100000000000 {
		task.RunTime = runtime
		runtime = runtime * 1000
	}

	return task
}

func (task *Task) SetID(uuid string) TaskSetInterface {
	task.ID = uuid
	return task
}

func (task *Task) SetSpacing(spacing int64) TaskSetInterface {
	task.Spacing = spacing
	return task
}

func (task *Task) SetEndTime(endTime int64) TaskSetInterface {
	task.EndTime = endTime
	return task
}

func (task *Task)  SetRunNumber(number int) TaskSetInterface {
	task.Number = number
	return task
}


func (task *Task) RunJob() {
	task.GetJob().Run()
}
func (task *Task) GetJob() Job {
	return task.Job
}
func (task *Task) GetUuid() string {
	return task.ID
}
func (task *Task) GetRunTime() int64 {
	return task.RunTime
}
func (task *Task) GetSpacing() int64 {
	return task.Spacing
}
func (task *Task) GetEndTime() int64 {
	return task.EndTime
}
func (task *Task) GetNumber() int {
	return task.Number
}