package utils

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IOSuite struct {
	suite.Suite
}

func (s *IOSuite) SetupTest() {
}

// TestMarshal :
func (s *IOSuite) TestHelloWorld() {
	BoWrite()
}
// TestMarshal :
func (s *IOSuite) TestOpeFile() {
	fmt.Println(MakeFileName("test", 1, "","use-",".py"))
}


func TestIOConfiguration(t *testing.T) {
	suite.Run(t, new(IOSuite))
}
