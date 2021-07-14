package goroutine_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/goroutine"
	"testing"
)

type GoroutineSuite struct {
	suite.Suite
}

func (s *GoroutineSuite) SetupTest() {
}

// TestMarshal :
func (s *GoroutineSuite) TestConfig() {
	for i := 0; i < 20; i++ {
		s.Equal(100, goroutine.UnsafeDemo())
	}

}

// TestMarshal :
func (s *GoroutineSuite) TestConfig2() {
	for i := 0; i < 20; i++ {
		s.Equal(100, goroutine.UnsafeDemo2())
	}

}

// TestMarshal :
func (s *GoroutineSuite) TestConfig3() {
	for i := 0; i < 20; i++ {
		s.Equal(100, goroutine.UnsafeDemo3())
	}
}
func (s *GoroutineSuite) TestConfig4() {

	goroutine.UnsafeDemo4()
}

// TestHttpClient :
func TestGoroutineSuite(t *testing.T) {
	suite.Run(t, new(GoroutineSuite))
}
