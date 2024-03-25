package utils

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BackOffSuite struct {
	suite.Suite
}

func (s *BackOffSuite) SetupTest() {
}

// TestMarshal :
func (s *BackOffSuite) TestHelloWorld() {
	// 带error 的重试
	count := 0
	// 重试3次 第四次成功
	s.NoError(ExponentialRetry(3, func() error {
		fmt.Println(fmt.Sprintf("重试%d次", count))
		count += 1
		if count == 4 {
			return nil
		}
		return errors.New("11")
	}))
	s.Equal(4, count)
}

func (s *BackOffSuite) TestDemo() {
	// 指数量级展示

}

func TestRetryConfiguration(t *testing.T) {
	suite.Run(t, new(BackOffSuite))
}
