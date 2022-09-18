package tasks

// Set方法

// SetJob 设置Job
func (task *Task) SetJob(job Job) TaskSetInterface {
	task.Job = job
	return task
}

// SetRunTime 设置runtime
func (task *Task) SetRunTime(runtime int64) TaskSetInterface {

	if runtime < 100000000000 {
		task.RunTime = runtime
		runtime = runtime * 1000
	}

	return task
}

// SetID 设置uuid
func (task *Task) SetID(uuid string) TaskSetInterface {
	task.ID = uuid
	return task
}

// SetSpacing 设置spacing
func (task *Task) SetSpacing(spacing int64) TaskSetInterface {
	task.Spacing = spacing
	return task
}

// SetEndTime 设置endTime
func (task *Task) SetEndTime(endTime int64) TaskSetInterface {
	task.EndTime = endTime
	return task
}

// SetRunNumber 设置number
func (task *Task)  SetRunNumber(number int) TaskSetInterface {
	task.Number = number
	return task
}


// RunJob 执行方法
func (task *Task) RunJob() {
	task.GetJob().Run()
}

// Get方法

// GetJob 取得Job
func (task *Task) GetJob() Job {
	return task.Job
}

// GetUUID 取得ID
func (task *Task) GetUUID() string {
	return task.ID
}

// GetRunTime 取得 RunTime
func (task *Task) GetRunTime() int64 {
	return task.RunTime
}

// GetSpacing 取得Spacing
func (task *Task) GetSpacing() int64 {
	return task.Spacing
}

// GetEndTime 取得EndTime
func (task *Task) GetEndTime() int64 {
	return task.EndTime
}

// GetNumber 取得Number
func (task *Task) GetNumber() int {
	return task.Number
}