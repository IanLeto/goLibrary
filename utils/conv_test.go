package utils_test

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type ConvSuite struct {
	suite.Suite
}

func (s *ConvSuite) SetupTest() {
}

// TestMarshal :
func (s *ConvSuite) TestHelloWorld() {
	// sql to arr

}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ConvSuite))
}
