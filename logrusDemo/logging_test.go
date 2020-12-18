package logrusDemo_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestLogSuit struct {
	suite.Suite
}

func (s *TestLogSuit) SetupTest() {

}

func (s *TestLogSuit) SimpleTest() {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestLogSuit))
}
