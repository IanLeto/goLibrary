package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
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


func TestRaSuite(t *testing.T) {
	suite.Run(t, new(TestRateSuit))
}
