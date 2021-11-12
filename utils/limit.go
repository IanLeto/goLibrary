package utils

import (
	"context"
	"sync"
)

type DemoTask struct {
}

func (d *DemoTask) Start() error {
	panic("implement me")
}

func (d *DemoTask) Work() error {
	panic("implement me")
}

func (d *DemoTask) Stop() error {
	panic("implement me")
}

func (d *DemoTask) Info(chan interface{}) {

}

func NewPool() {
	//limiter := rate.NewLimiter(10, )
}

type Manager struct {
	wg    sync.WaitGroup
	ctx   context.Context
	queue []*Task
}

func (m Manager) Receive(t []DemoTask, info chan interface{}) {

}
