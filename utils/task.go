package utils

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

type Task interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
}

type EchoRandTask struct {
	Rank int
	ID   int
}

func (e *EchoRandTask) Run(ctx context.Context) error {
	time.Sleep(1 * time.Second)
	fmt.Println(e.ID, ":", rand.Intn(1000))
	return nil
}

func (e *EchoRandTask) Stop(ctx context.Context) error {
	panic("implement me")
}

type CallBackFuncType func(err error)

type TaskManager struct {
	// 待完成的任务队列, 外部可以不断往里推送任务
	JobQueue chan Task

	// worker并发任务的令牌桶
	tokenBucket chan interface{}
	maxWorker   int
	// 任务结束标志位
	ctx        context.Context
	cancelFunc context.CancelFunc
	// 回调函数的方法
	callBackFunc CallBackFuncType
	// 等待所有任务的结束的waitGroup
	wg sync.WaitGroup
	// 任务是否正在进行的标记位
	isRunning bool
	// 所有任务列表
	taskList  []Task
	jobWg     sync.WaitGroup
	startTime time.Time
	endTime   time.Time
}

func (m *TaskManager) PushTask() {
	// 开始执行任务
	m.startTime = time.Now()
	go func() {
		defer func() {
			// 任务关停
			m.isRunning = false
		}()
		for {
			select {
			case task := <-m.JobQueue:
				// 当我们拿到一个任务的时候，先去看看是否有可用的执行令牌
				<-m.tokenBucket
				// 拿到执行令牌，开始派遣goroutine 执行
				m.wg.Add(1)
				go func(job Task) {
					defer func() {
						// 本case 执行结束
						m.wg.Done()
						// 交还令牌
						m.tokenBucket <- struct{}{}
						// 总任务量
						m.jobWg.Done()
					}()
					if err := job.Run(m.ctx); err != nil {
						return
					}
					//m.callBackFunc(job.Run(m.ctx))
				}(task)
			case <-m.ctx.Done():
				return
			}
		}
	}()

	go func() {
		defer func() {
			m.jobWg.Done()
		}()
		for _, task := range m.taskList {
			select {
			case m.JobQueue <- task:
				m.jobWg.Add(1)
			// 考虑到有缓冲chan 关闭问题，这里选择使用ctx 主动关闭任务推送
			case <-m.ctx.Done():
				return
			}

		}
	}()
	//我们认为，push 任务到chan 也是需要时间的
	m.jobWg.Add(1)

}

func (m *TaskManager) WaitJob() error {
	m.jobWg.Wait()
	// 关闭所有任务
	return nil
}
func (m *TaskManager) Wait() error {
	m.wg.Wait()
	// 关闭所有任务
	return nil
}

func (m *TaskManager) Stop() error {
	m.cancelFunc()
	m.endTime = time.Now()
	logrus.Debugf("耗时%ds", m.endTime.Sub(m.startTime))
	return nil
}
func NewTaskManager(ctx context.Context, cancelFunc context.CancelFunc, taskList []Task,
	jobQueue chan Task, tokenBucket chan interface{}, callBack CallBackFuncType) *TaskManager {
	return &TaskManager{
		ctx:          ctx,
		cancelFunc:   cancelFunc,
		taskList:     taskList,
		JobQueue:     jobQueue,
		tokenBucket:  tokenBucket,
		callBackFunc: callBack,
		jobWg:        sync.WaitGroup{},
		wg:           sync.WaitGroup{},
	}
}
