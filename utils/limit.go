package utils

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

type DemoTask struct {
}

func (d *DemoTask) Start() error {
	time.Sleep(1 * time.Millisecond)
	println(rand.Intn(1000))
	return nil
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
	wg  sync.WaitGroup
	ctx context.Context
}

func (m Manager) Receive(t []DemoTask, info chan interface{}) {

}
