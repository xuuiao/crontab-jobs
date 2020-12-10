package main

import (
	"fmt"
	"github.com/xuuiao/crontab-jobs/cronjob"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	job := cronjob.NewJob()
	job.Start()
	ObserveExitSignal(func(o os.Signal) {
		if err := job.ShutDown(); err != nil {
			fmt.Printf("job shutdown failed, err:%s \n", err.Error())
		}
	})
}

// ObserveExitSignal 监听系统退出信号
// 会阻塞当前运行
func ObserveExitSignal(f func(os.Signal)) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("get signal %s \n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			f(s)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
