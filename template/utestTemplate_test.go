package template_test

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type TemplateSuite struct {
	suite.Suite
}

func (s *TemplateSuite) SetupTest() {
}

// TestMarshal :
func (s *TemplateSuite) TestHelloWorld() {
}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(TemplateSuite))
}
