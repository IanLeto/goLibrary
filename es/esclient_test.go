package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/suite"
	"testing"
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

func TestES(t *testing.T) {
	suite.Run(t, new(TestESSuit))
}
