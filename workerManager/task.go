package workerManager

import "context"

type Task struct {
	HostList []string
}

func NewTaskManager(ctx context.Context, maxWorker int) {

}
