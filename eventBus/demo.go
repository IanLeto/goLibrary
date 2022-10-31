package eventBus

import (
	"fmt"
	"goLibrary/utils"
	"time"
)

type DemoTask struct {
}

// step 2 run func
func (t *DemoTask) Run() {
	var (
		// step3 声明一个chan 用来接受外部事件
		ch = make(chan string, 0)
	)
	ticker := time.NewTicker(5 * time.Second)
	// step4 定义外来事件的处理方式
	err := GlobalEvent.SubscribeAsync("demoInsert", func(id string) {
		ch <- id
	}, false)
	utils.NoErr(err)
	for {
		select {
		case <-ticker.C:
			go func() {
				time.Sleep(10 * time.Second)
				fmt.Println("周期任务", time.Now())
			}()
		case id := <-ch:
			go func(id string) {
				fmt.Println("插队任务", id)
			}(id)
		}

	}
}

// step1 初始化一个周期任务
func NewTask() *DemoTask {
	return &DemoTask{}
}
