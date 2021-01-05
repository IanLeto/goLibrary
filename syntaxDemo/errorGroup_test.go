package syntaxDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/syntaxDemo"
	"testing"
)

type TestErrorGroupSuit struct {
	suite.Suite
}

func (s *TestErrorGroupSuit) SetupTest() {

}

func (s *TestErrorGroupSuit) TestSimpleTest() {
	syntaxDemo.ErrorGroupDemoRun()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestErrorGroupSuit))
}
