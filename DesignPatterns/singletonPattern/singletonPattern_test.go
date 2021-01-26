package singletonPattern_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/DesignPatterns/singletonPattern"
	"sync"
	"testing"
)

type TestSingletonPatternSuit struct {
	suite.Suite
}

func (s *TestSingletonPatternSuit) SetupTest() {
	instance1 := singletonPattern.GetInstance()
	instance2 := singletonPattern.GetInstance()
	s.Equal(instance1, instance2)
	instance3 := singletonPattern.GetInstanceNotOnce()
	instance4 := singletonPattern.GetInstanceNotOnce()
	s.Equal(instance3, instance4)
}
// 该测试的目的为每次生成的instance 的指针都是相同的 因此，无论多少次调用 getinstace 都会生成相同的实例
func (s *TestSingletonPatternSuit) TestSimpleTest() {
	var (
		wg    = sync.WaitGroup{}
		start = make(chan struct{})
	)

	wg.Add(100)
	instances := [100]*singletonPattern.Singleton{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			<-start
			instances[index] = singletonPattern.GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < 100; i++ {
		s.Equal(instances[i], instances[i-1])
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSingletonPatternSuit))
}
