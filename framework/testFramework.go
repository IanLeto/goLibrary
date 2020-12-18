package framework

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestDemoSuit struct {
	suite.Suite
}

func (s *TestDemoSuit) SetupTest() {

}


func TestSuite(t *testing.T) {
	suite.Run(t, new(TestDemoSuit))
}
