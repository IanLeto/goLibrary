package prometheusDemo_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/prometheusDemo"
	"testing"
)

type TestPromClientSuit struct {
	suite.Suite
}

func (s *TestPromClientSuit) SetupTest() {

}

func (s *TestPromClientSuit) TestPush() {
	prometheusDemo.Pusher()
}

func (s *TestPromClientSuit) TestSSTest() {
	prometheusDemo.ValueModel()
}

func (s *TestPromClientSuit) TestSimpleTest() {
	prometheusDemo.HelloWorld()
	//prometheusDemo.PromQueryRequest("http://172.20.178.124:9091","up")
	//var res = make(chan int)
	//var x = make([]int, 0, 1000)
	//var sg sync.WaitGroup
	//sg.Add(1000)
	//for i := 0; i < 1000; i++ {
	//	go func(ss int) {
	//		defer sg.Done()
	//		res <- ss
	//	}(i)
	//}
	//go func() {
	//	for i := range res {
	//		fmt.Println(i)
	//		x = append(x, i)
	//	}
	//}()
	//sg.Wait()
	//time.Sleep(1 * time.Second)
	//close(res)
	//fmt.Println(x)

}

func TestPromClientSuite(t *testing.T) {
	suite.Run(t, new(TestPromClientSuit))
}
