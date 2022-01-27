package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"goLibrary/httpServerInDocker"
	"goLibrary/utils"
	"testing"
)

type TestRateSuit struct {
	suite.Suite
}

func (s *TestRateSuit) SetupTest() {
	//var base float64 = 100000
	//var rate float64 = 6.8 / base
	//for i := 0; i < 365*60; i++ {
	//	base = base * (1 + rate)
	//}
	//fmt.Println(base)
}

// 银行
func (s *TestRateSuit) TestRate() {
	var base float64 = 100000
	var rate float64 = 0.036
	for i := 0; i < 20; i++ {
		base = base * (1 + rate)
	}
	fmt.Println(base)
}

// 通胀
func (s *TestRateSuit) TestRate2() {
	var base float64 = 100000
	var rate float64 = 0.036
	for i := 0; i < 20; i++ {
		base = base * (1 - rate)
	}
	fmt.Println(base)
}

func (s *TestRateSuit) TestEnv() {
	s.Equal(utils.GetLocalOSEnv("CCMODE"), "DEBUG")
}

func (s *TestRateSuit) TestFastDemo() {
	httpServerInDocker.FastHttpDemo()
}
func (s *TestRateSuit) TestBatch() {
	s.Equal([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}}, utils.Batch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
}

func TestRaSuite(t *testing.T) {
	suite.Run(t, new(TestRateSuit))
}
