package cron

import (
	"crontask/tasks"
	"github.com/google/uuid"
	"time"
)

// AddFuncSpaceNumber 加入调度对列中
func (scheduler *TaskScheduler) AddFuncSpaceNumber(spaceTime int64, number int, f func())  {
	task := tasks.GetTaskWithFuncSpacingNumber(spaceTime, number, f)
	_ = scheduler.AddTask(task)
}

// AddFuncSpace 加入调度对列中
func (scheduler *TaskScheduler) AddFuncSpace(spaceTime int64, endTime int64, f func())  {
	task := tasks.GetTaskWithFuncSpacing(spaceTime, endTime, f)
	_ = scheduler.AddTask(task)
}

// AddFunc 加入调度对列中
func (scheduler *TaskScheduler) AddFunc(unixTime int64, f func())  {
	task := tasks.GetTaskWithFunc(unixTime, f)
	_ = scheduler.AddTask(task)
}

// AddTask 加入调度对列中
func (scheduler *TaskScheduler) AddTask(task *tasks.Task) string {

	// 筛选条件。
	if task.RunTime < 100000000000 {
		task.RunTime = task.RunTime * int64(time.Second)
	}

	if task.RunTime < time.Now().UnixNano() {
		task.RunTime = time.Now().UnixNano() + int64(time.Second)
	}

	if task.ID == "" {
		task.ID = uuid.New().String()
	}

	return scheduler.addTask(task)

}


func (scheduler *TaskScheduler) addTask(task tasks.TaskInterface) string {
	if scheduler.lock {
		scheduler.swap = append(scheduler.swap, task)
	} else {
		scheduler.tasks = append(scheduler.tasks, task)
		scheduler.add <-task
	}

	return task.GetUUID()

}

// ExportInterface 调用task实现的接口方法
func (scheduler *TaskScheduler) ExportInterface() []tasks.TaskInterface {
	return scheduler.tasks
}
// Export 查看task细节
func (scheduler *TaskScheduler) Export() []*tasks.Task {
	task := make([]*tasks.Task,0)
	for _,v := range scheduler.tasks {
		task = append(task, v.(*tasks.Task))
	}
	return task
}

// StopOnce 删除
func (scheduler *TaskScheduler) StopOnce(uuidStr string) {
	scheduler.remove <- uuidStr
}

// Start 启动调度对列
func (scheduler *TaskScheduler) Start() {

	go scheduler.run()
}

// Stop 停止调度对列工作
func (scheduler *TaskScheduler) Stop() {
	scheduler.stop <- struct{}{}
}

//run tasks list
//if is empty, run a year timer tasks
func (scheduler *TaskScheduler) run() {

	for {

		now := time.Now()
		task, key := scheduler.GetTask()
		//0 阻断
		if key < 0 {
			continue
		}
		runTime := task.GetRunTime()
		i64 := runTime - now.UnixNano()

		var d time.Duration
		if i64 < 0 {
			// 设置时间
			scheduler.tasks[key].SetRunTime(now.UnixNano())
			if task != nil {
				go task.RunJob()
			}
			scheduler.doAndReset(key)
			continue
		} else {
			sec := runTime / int64(time.Second)
			nsec := runTime % int64(time.Second)

			d = time.Unix(sec, nsec).Sub(now)
		}

		timer := time.NewTimer(d)

		//catch a chan and do something
		for {
			select {
			//if time has expired do tasks and shift key if is tasks list
			case <-timer.C:
				scheduler.doAndReset(key)
				if task != nil {
					//fmt.Println(scheduler.tasks[key])
					go task.RunJob()
					timer.Stop()
				}

				//if add tasks
			case <-scheduler.add:
				timer.Stop()
				// remove tasks with remove uuid
			case uuidstr := <-scheduler.remove:
				scheduler.removeTask(uuidstr)
				timer.Stop()
				//if get a stop single exit
			case <-scheduler.stop:
				timer.Stop()
				return
			}

			break
		}
	}
}


// GetTask 取得task对象
func (scheduler *TaskScheduler) GetTask() (task tasks.TaskGetInterface, tempKey int) {
	scheduler.Lock()
	defer scheduler.UnLock()

	// 没有task
	if len(scheduler.tasks) < 1 {
		return nil,-1
	}

	min := scheduler.tasks[0].GetRunTime()
	tempKey = 0

	for key, task := range scheduler.tasks {
		tTime := task.GetRunTime()
		if min <= tTime {
			continue
		}
		if min > tTime {
			tempKey = key

			min = tTime
			continue
		}
	}

	task = scheduler.tasks[tempKey]

	return task, tempKey
}

//if add a new tasks and runtime < now tasks runtime
// stop now timer and again
func (scheduler *TaskScheduler) doAndReset(key int) {
	scheduler.Lock()
	defer scheduler.UnLock()
	//null pointer
	if key < len(scheduler.tasks) {

		nowTask := scheduler.tasks[key]
		scheduler.tasks = append(scheduler.tasks[:key], scheduler.tasks[key+1:]...)

		if nowTask.GetSpacing() > 0 {
			tTime := nowTask.GetRunTime()
			nowTask.SetRunTime(nowTask.GetSpacing() * int64(time.Second) + tTime)
			number := nowTask.GetNumber()
			if number > 1 {
				nowTask.SetRunNumber(number - 1)
				scheduler.tasks = append(scheduler.tasks, nowTask)
			} else if nowTask.GetEndTime() >= tTime {
				scheduler.tasks = append(scheduler.tasks, nowTask)
			}
		}

	}
}


// 由ID 找的task并删除
func (scheduler *TaskScheduler) removeTask(uuidStr string) {
	scheduler.Lock()
	defer scheduler.UnLock()
	for key, task := range scheduler.tasks {
		if task.GetUUID() == uuidStr {
			scheduler.tasks = append(scheduler.tasks[:key], scheduler.tasks[key+1:]...)
			break
		}
	}
}


// Lock 加锁操作
func (scheduler *TaskScheduler) Lock() {
	scheduler.lock = true
}

// UnLock 解锁操作
func (scheduler *TaskScheduler) UnLock() {
	scheduler.lock = false
	// 解锁后，要检查 swap对列中有没有task，有需要加入调度对列，并设置为0
	if len(scheduler.swap) > 0 {
		for _, task := range scheduler.swap {
			scheduler.tasks = append(scheduler.tasks, task)
		}
		scheduler.swap = make([]tasks.TaskInterface, 0)
	}
}
