package task

import (
	"context"
	"fmt"
	"goLibrary/utils"
	"time"
)

type DemoTask struct {
	ctx context.Context
}

func (d *DemoTask) Work() error {
	fmt.Println("demo")
	return nil
}

func NewDemoTask() *DemoTask {
	return &DemoTask{
		ctx: context.Background(),
	}

}

func (d *DemoTask) Start() error {
	utils.NoErr(NewCronTask(d.ctx, d, *time.NewTicker(5 * time.Second)))
	return nil
}

func (d DemoTask) Stop() error {
	panic("implement me")
}
