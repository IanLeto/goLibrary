package utils_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CMDSuite struct {
	suite.Suite
}

func (s *CMDSuite) SetupTest() {
}

// TestMarshal :
func (s *CMDSuite) TestHelloWorld() {

}

func TestCmd(t *testing.T) {
	suite.Run(t, new(CMDSuite))
}
