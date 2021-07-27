package base

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SelectSuite struct {
	suite.Suite
}

func (s *SelectSuite) SetupTest() {
}

// TestMarshal :
func (s *SelectSuite) TestHelloWorld() {
	RunSelect(nil)
}

func TestRunSelect(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}
