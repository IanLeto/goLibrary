package logDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/logDemo"
	"testing"
)

type TestZapLogSuit struct {
	suite.Suite
}

func (s *TestZapLogSuit) SetupTest() {
	logDemo.InitLogger()
}

func (s *TestZapLogSuit) TestConf() {
	logDemo.Logger.Error("x")

}

func (s *TestZapLogSuit) TestHook() {
	//logDemo.UseHook()

}

func TestZapLog(t *testing.T) {
	suite.Run(t, new(TestZapLogSuit))
}
