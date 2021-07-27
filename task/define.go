package task

import "fmt"

var TaskQueue = make([]Task, 0)

type Task interface {
	Start() error
	Work() error
	Stop() error
}

type CacheTask struct {
}

func (c CacheTask) Start() error {
	fmt.Println(1)
	return nil
}

func (c CacheTask) Stop() error {
	panic("implement me")
}
