package logrusDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/logrusDemo"
	"testing"
)

type TestLogSuit struct {
	suite.Suite
}

func (s *TestLogSuit) SetupTest() {

}

func (s *TestLogSuit) TestConf() {
	//logrusDemo.NormalLog()
	//logrusDemo.SetLogFormatter()

}

func (s *TestLogSuit) TestHook() {
	//logrusDemo.UseHook()
	logrusDemo.UseDivLogFile()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestLogSuit))
}
