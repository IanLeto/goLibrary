package config_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"goLibrary/config"
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
	s.Equal("debug", config.BaseConfig.RunMode)
	var worker = make(chan interface{}, 50)
	var tokenBullet = make(chan interface{}, 50)
	var jobQuene = make(chan int)
	for {
		select {
		case v, ok := <-jobQuene:
			if !ok {
				return
			}
			go func(s int) {

			}(v)


		}

	}
}

// TestViperConfiguration :
func TestConfiguration(t *testing.T) {
	suite.Run(t, new(TestInitConfigSuit))
}
