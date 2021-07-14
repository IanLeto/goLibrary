package goroutine_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/goroutine"

	"testing"
)

type TemplateSuite struct {
	suite.Suite
}

func (s *TemplateSuite) SetupTest() {

}

// TestMarshal :
func (s *TemplateSuite) TestHowToCount() {
	goroutine.HowToCount()
}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(TemplateSuite))
}
