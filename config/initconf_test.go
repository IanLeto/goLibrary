package config_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	_ "goLibrary/config"
	"goLibrary/testsuites"
	"sync"
	"testing"
)

// TestInitConfigSuit :
type TestInitConfigSuit struct {
	suite.Suite
	testsuites.ConfigSuit
}

func (s *TestInitConfigSuit) SetupTest() {
	s.InitConfigSuit()
}

type taskManager struct {
	maxWorker   int
	jobQueue    chan string
	tokenBucket chan interface{}
	ctx         context.Context
	jobWg       sync.WaitGroup
}

func (s *taskManager) start() {
	for {
		select {
		case <-s.jobQueue:
			go func() {

			}()

		}
	}
}

// TestMarshal :
func (s *TestInitConfigSuit) TestConfig() {

}

// TestViperConfiguration :
func TestConfiguration(t *testing.T) {
	suite.Run(t, new(TestInitConfigSuit))
}
