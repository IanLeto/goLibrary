package viperDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/viperDemo"

	"testing"
)

type ViperDemoSuite struct {
	suite.Suite
}

func (s *ViperDemoSuite) SetupTest() {
}

func (s *ViperDemoSuite) TestHelloWorld() {
	viperDemo.RunEcho()
}

func TestViperDemoConfiguration(t *testing.T) {
	suite.Run(t, new(ViperDemoSuite))
}
