package pprof_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/pprof"

	"testing"
)

type PToolsSuite struct {
	suite.Suite
}

func (s *PToolsSuite) SetupTest() {
}

// TestMarshal :
func (s *PToolsSuite) TestHelloWorld() {
	go func() {
		for {
			pprof.BeTest()
		}
	}()
	pprof.PprofWeb()

}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(PToolsSuite))
}
