package crontask

import (
	cron2 "crontask/cron"
	"crontask/tasks"
	"fmt"
	"testing"
	"time"
)

func Test_AddFunc(t *testing.T) {
	cron := cron2.NewCron()

	go cron.Start()

	cron.AddFunc(time.Now().UnixNano()+int64(time.Second), func() {
		fmt.Println("one second after")
	})
	cron.Logger.Println(fmt.Println("log out!!"))
	cron.AddFunc(time.Now().UnixNano()+int64(time.Second * 3), func() {
		fmt.Println("three second after")
	})

	cron.AddFunc(time.Now().UnixNano()+int64(time.Second * 10), func() {
		fmt.Println("ten second after")
	})

	taskl := cron.ExportInterface()
	for _, task := range taskl {
		fmt.Println(task.GetUUID())
	}


	tasklist := cron.Export()
	for _, task := range tasklist {
		fmt.Println(task.ID)
	}

	timer := time.NewTimer(20 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("任务结束")
			return
		default:
		}
	}

}

//test add space task func
func Test_AddFuncSpace(t *testing.T) {
	cron := cron2.NewScheduler()

	go cron.Start()

	cron.AddFuncSpace(1, time.Now().UnixNano()+int64(time.Second*1), func() {
		fmt.Println("one second after")
	})

	//cron.AddFuncSpace(1, time.Now().UnixNano()+int64(time.Second*20), func() {
	//	fmt.Println("one second after, task second")
	//})

	cron.AddFunc(time.Now().UnixNano()+int64(time.Second*10), func() {
		fmt.Println("ten second after")
	})
	timer := time.NewTimer(11 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("over")
		}
		break
	}

}

//test add Task and timing add Task
func Test_AddTask(t *testing.T) {
	cron := cron2.NewCron()
	go cron.Start()

	cron.AddTask(&tasks.Task{
		Job:tasks.FuncJob(func() {
			fmt.Println("hello cron")
		}),
		RunTime:time.Now().UnixNano()+int64(time.Second*2),
	})


	cron.AddTask(&tasks.Task{
		Job:tasks.FuncJob(func() {
			fmt.Println("hello cron1")
		}),
		RunTime:time.Now().UnixNano()+int64(time.Second*3),
	})

	//cron.AddTask(&tasks.Task{
	//	Job: tasks.FuncJob(func() {
	//		fmt.Println("hello cron2 loop")
	//	}),
	//	RunTime: time.Now().UnixNano() + int64(time.Second*4),
	//	Spacing: 1,
	//	EndTime: time.Now().UnixNano() + 9*(int64(time.Second)),
	//})

	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("over")
		}
		break
	}
}


