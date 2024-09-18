package codingBucket_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BucketSuite struct {
	suite.Suite
}

func (s *BucketSuite) SetupTest() {

}

func (s *BucketSuite) TestSimpleJob() {

}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(BucketSuite))
}
