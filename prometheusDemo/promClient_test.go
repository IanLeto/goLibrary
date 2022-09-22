package prometheusDemo_test

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/suite"
	"goLibrary/prometheusDemo"
	"testing"
	"time"
)

type TestPromClientSuit struct {
	suite.Suite
	api v1.API
	ctx context.Context
}

func (s *TestPromClientSuit) SetupTest() {
	cli, _ := api.NewClient(api.Config{
		Address: "http://124.222.48.125:9090",
	})
	s.api = v1.NewAPI(cli)
	s.ctx = context.Background()
}

func (s *TestPromClientSuit) TestPush() {
	prometheusDemo.Pusher()
}

func (s *TestPromClientSuit) TestSSTest() {
}

// 基础cpu 使用率
func (s *TestPromClientSuit) TestQuery1() {
	// 10m 之前
	start := time.Now().Add(time.Duration(-10) * time.Minute)

	res, err := s.api.QueryRange(s.ctx, "sum without(cpu, mode) (node_cpu_seconds_total)", v1.Range{
		Start: start,
		End:   time.Now(),
		Step:  time.Duration(30) * time.Second,
	})
	s.NoError(err)
	s.NoError(err)
	v, ok := res.(model.Matrix)
	s.Equal(true, ok)
	for _, stream := range v {
		fmt.Println(stream.Metric, stream.Values)
	}
}

//3262639057
func BenchmarkQuery(b *testing.B) {

	cli, _ := api.NewClient(api.Config{
		Address: "http://124.222.48.125:9090",
	})
	api := v1.NewAPI(cli)
	for i := 0; i < b.N; i++ {
		_, err := api.Query(context.TODO(), "go_gc_duration_seconds[14d]", time.Now())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

//90230955
//92693207
//82626286
//139188935
//121219658
func BenchmarkQueryRange(b *testing.B) {

	cli, _ := api.NewClient(api.Config{
		Address: "http://124.222.48.125:9090",
	})
	api := v1.NewAPI(cli)
	start := time.Now().Add(time.Duration(-24) * time.Hour * 7 * 2)

	for i := 0; i < b.N; i++ {
		_, err := api.QueryRange(context.TODO(), "node_entropy_available_bits", v1.Range{
			Start: start,
			End:   time.Now(),
			Step:  time.Duration(24) * time.Hour,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (s *TestPromClientSuit) TestSimpleTest() {
	//prometheusDemo.HelloWorld()
	prometheusDemo.NewPromMetrics()
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
