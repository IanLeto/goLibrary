package logDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/logDemo"
	"testing"
)

type TestLogSuit struct {
	suite.Suite
}

func (s *TestLogSuit) SetupTest() {

}

func (s *TestLogSuit) TestConf() {

}

func (s *TestLogSuit) TestHook() {
	//logDemo.UseHook()
	logDemo.UseDivLogFile()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestLogSuit))
}
