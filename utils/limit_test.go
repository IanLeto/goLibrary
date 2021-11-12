package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestListSuit struct {
	suite.Suite
}

func (s *TestListSuit) SetupTest() {
	type config struct {
		Max      int // 当前最大进程
		CallBack func(chan error)
	}

}

// 银行
func (s *TestListSuit) TestList() {

	for _, i2 := range []int{1, 2, 3} {
		if i2 > 2 {
			fmt.Println(111)
			continue
		}
		fmt.Println(22222)
	}

}

func (s *TestListSuit) TestFastDemo() {
}

func TestLimitSuite(t *testing.T) {
	suite.Run(t, new(TestListSuit))
}
