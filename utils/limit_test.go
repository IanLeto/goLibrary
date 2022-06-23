package utils_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestListSuit struct {
	suite.Suite
}

func (s *TestListSuit) SetupTest() {

}

func (s *TestListSuit) TestFastDemo() {

}

func TestLimitSuite(t *testing.T) {
	suite.Run(t, new(TestListSuit))
}
