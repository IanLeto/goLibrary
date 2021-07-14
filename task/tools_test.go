package task_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/task"

	"testing"
)

type TaskSuite struct {
	suite.Suite
	task task.Task
}

func (s *TaskSuite) SetupTest() {
}

// TestMarshal :
func (s *TaskSuite) TestHelloWorld() {


}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(TaskSuite))
}
