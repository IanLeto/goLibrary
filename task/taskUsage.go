package task

import (
	"goLibrary/utils"
)

func InitTask() {
	for _, task := range TaskQueue {
		utils.NoErr(task.Start())
	}
}

func init() {
	TaskQueue = append(TaskQueue, NewDemoTask())

}
