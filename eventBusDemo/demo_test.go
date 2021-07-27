package eventBusDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/eventBusDemo"
	"testing"
	"time"
)

type SubPubSuite struct {
	suite.Suite
}

func (s *SubPubSuite) SetupTest() {
}

// 如何使用
func (s *SubPubSuite) TestHelloWorld() {
	// step 1 初始化并运行
	t := eventBusDemo.NewTask()
	// step 2 抽象出一个插入事件/aka 发布事件
	go func() {
		t := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-t.C:
				// step 3 发布事件以及参数
				eventBusDemo.GlobalEvent.Publish("demoInsert", "1")
			}

		}
	}()
	t.Run()

}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(SubPubSuite))
}
