package task

import "fmt"

type Task interface {
	Start() error
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

