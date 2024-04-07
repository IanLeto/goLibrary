package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CodeReview struct {
	suite.Suite
}

func (s *CodeReview) SetupTest() {

}

func (s *CodeReview) TestMySQL() {

}

func TestCodeReviewSuite(t *testing.T) {
	suite.Run(t, new(CodeReview))

}
