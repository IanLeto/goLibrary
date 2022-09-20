package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/suite"
	"goLibrary/utils"
	"math/rand"
	"testing"
	"time"
)

type TestESSuit struct {
	suite.Suite
	session *mgo.Session
	client  *elastic.Client
	ctx     context.Context
}

func (s *TestESSuit) SetupTest() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://124.222.48.125:9200/"))
	s.NoError(err)
	s.client = client
	s.ctx = context.Background()

}

type Sp struct {
	ID      int
	Name    string
	BigText string
	Script  string
	AAA     string
	BBB     string
}

// 长度为62
var bytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func init() {
	// 保证每次生成的随机数不一样
	rand.Seed(time.Now().UnixNano())
}

// 方法一
func RandStr1(n int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Int31()%62]
	}
	return string(result)
}

func (s *TestESSuit) TestCreateForBench() {
	var ch = make(chan struct{}, 10)
	for i := 0; i < 200000; i++ {
		i := i
		ch <- struct{}{}
		go func(chan struct{}) {
			_, _ = s.client.Index().Index("benchmark").BodyJson(&Sp{
				ID:      i,
				Name:    fmt.Sprintf("name%d", i),
				BigText: RandStr1(40),
				Script:  fmt.Sprintf("script%s", cast.ToString(i)),
				AAA:     "",
				BBB:     "",
			}).Do(context.Background())
			defer func() {
				<-ch
			}()
		}(ch)

	}

}

func (s *TestESSuit) TestSimpleTest() {
	var body = struct {
		Name string `json:"name"`
	}{
		Name: "ianleto",
	}

	res, err := s.client.Index().Index("get-x").BodyJson(body).Do(context.Background())
	s.NoError(err)
	r, _ := json.Marshal(res)
	fmt.Println(string(r))
	fmt.Println(body.Name)

	r2, err := s.client.Get().Index("get-x").Do(context.Background())
	s.NoError(err)
	fmt.Println(r2)
}

func formatHit(result *elastic.SearchResult) {
	defer utils.TimeCost()()
	for _, hit := range result.Hits.Hits {
		r, _ := hit.Source.MarshalJSON()
		fmt.Println(*hit.Score, ":", string(r))
	}
}

func (s *TestESSuit) TestQuery() {
	bq := elastic.NewBoolQuery()
	bq.MustNot(elastic.NewTermQuery("Uploader", ""))
	// 过滤字段
	eq := elastic.NewExistsQuery("Uploader")
	res, err := s.client.Search().Index("script").Query(bq).Query(eq).Do(s.ctx)
	s.NoError(err)
	formatHit(res)
}

func (s *TestESSuit) TestDel() {
	bq := elastic.NewBoolQuery()
	bq.Must(elastic.NewTermQuery("Content", "echo"))
	res, err := s.client.DeleteByQuery().Index("script").Query(bq).Do(s.ctx)
	s.NoError(err)
	fmt.Println(res.Total)
	//formatHit(res)
}

// 基础查询，查询 uploader 为 ian 且有效的数据
func (s *TestESSuit) TestQueryBase() {
	bq := elastic.NewBoolQuery()
	bq.Filter(elastic.NewTermQuery("Uploader", "ian"))
	eq := elastic.NewExistsQuery("Uploader")
	res, err := s.client.Search().Index("script").Query(bq).Query(eq).Do(s.ctx)
	s.NoError(err)
	formatHit(res)

}

// 基础查询，查询 脚本中 有docker 相关的数据 且name中不含有test
func (s *TestESSuit) TestQuery2() {
	bq := elastic.NewBoolQuery()
	bq.Filter(elastic.NewTermQuery("Uploader", "ian"))
	bq.Must(elastic.NewMatchQuery("Content", "docker"))
	bq.MustNot(elastic.NewWildcardQuery("Name", "test*"))
	eq := elastic.NewExistsQuery("Uploader")
	// 注意顺序
	res, err := s.client.Search().Index("script").Query(eq).Query(bq).Do(s.ctx)
	s.NoError(err)
	formatHit(res)
}

// 基础查询3，查询 脚本中 基础id
func (s *TestESSuit) TestQuery3() {
	bq := elastic.NewBoolQuery()
	bq.Filter(elastic.NewTermQuery("_id", "Q1TzOYMBtW3yHWAAjPx3"))
	// 注意顺序
	res, err := s.client.Search().Index("script").Query(bq).Do(s.ctx)
	s.NoError(err)
	formatHit(res)
}

// base
func BenchmarkQuery(b *testing.B) {
	client, _ := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://124.222.48.125:9200/"))
	bq := elastic.NewBoolQuery()
	bq.Must(elastic.NewMatchQuery("ID", 100))
	for i := 0; i < b.N; i++ {
		_, _ = client.Search().Index("benchmark").Query(bq).Do(context.TODO())
	}

}

// 多线程
func BenchmarkQuery2(b *testing.B) {
	var ch = make(chan struct{}, 10)
	client, _ := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://124.222.48.125:9200/"))
	bq := elastic.NewBoolQuery()
	bq.Filter(elastic.NewMatchQuery("ID", 100))
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
		go func(chan struct{}) {
			_, _ = client.Search().Index("benchmark").Query(bq).Do(context.TODO())
		}(ch)
		<-ch
	}
}

func TestES(t *testing.T) {
	suite.Run(t, new(TestESSuit))
}
