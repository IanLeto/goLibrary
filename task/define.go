package task

type Task interface {
	Start() error
	Stop() error
}

type CacheTask struct {
}

func (c CacheTask) Start() error {
	panic("implement me")
}

func (c CacheTask) Stop() error {
	panic("implement me")
}
