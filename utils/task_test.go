package utils_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"goLibrary/utils"
	"testing"
)

// HttpClientSuite :
type TaskSuite struct {
	suite.Suite
	*utils.DemoTask
}

func (s *TaskSuite) SetupTest() {

}

func (s *TaskSuite) TestTaskManager() {
	var (
		ctx           = context.Background()
		jobQueue      = make(chan utils.Task, 50)
		tokenBucket   = make(chan interface{}, 50)
		taskList      []utils.Task
		idCtx, cancel = context.WithCancel(ctx)
	)
	for i := 0; i < 1000; i++ {
		taskList = append(taskList, &utils.EchoRandTask{
			Rank: i,
			ID:   i,
		})
	}
	for i := 0; i < 50; i++ {
		tokenBucket <- struct{}{}
	}
	var master = utils.NewTaskManager(idCtx, cancel, taskList, jobQueue, tokenBucket, nil)
	master.PushTask()
	s.NoError(master.Wait())
	s.NoError(master.WaitJob())
	s.NoError(master.Stop())
}

// 有回调版本
func (s *TaskSuite) TestTaskManagerCallBack() {
	var (
		ctx           = context.Background()
		jobQueue      = make(chan utils.Task, 50)
		tokenBucket   = make(chan interface{}, 50)
		taskList      []utils.Task
		idCtx, cancel = context.WithCancel(ctx)
	)
	for i := 0; i < 1000; i++ {
		taskList = append(taskList, &utils.EchoRandTask{
			Rank: i,
			ID:   i,
		})
	}
	for i := 0; i < 50; i++ {
		tokenBucket <- struct{}{}
	}

	// 使用有回调版本可以定制
	var fn utils.CallBackFuncType
	fn = func(t utils.Task) {
		eTask, ok := t.(*utils.EchoRandTask)
		if !ok {
			return
		}
		eTask.ID += 100
		s.NoError(eTask.Run(idCtx))
	}
	var master = utils.NewTaskManager(idCtx, cancel, taskList, jobQueue, tokenBucket, fn)
	master.PushTaskCallback()
	s.NoError(master.Wait())
	s.NoError(master.WaitJob())
	s.NoError(master.Stop())
}

func TestDashboard(t *testing.T) {
	suite.Run(t, new(TaskSuite))
}
