package goroutine_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/goroutine"
	"testing"
)

type InterfaceSuite struct {
	suite.Suite
}

func (s *InterfaceSuite) SetupTest() {

}

func (s *InterfaceSuite) TestHowToCount() {
	goroutine.TimeoutModel()

}

func TestInterface(t *testing.T) {
	suite.Run(t, new(InterfaceSuite))
}
