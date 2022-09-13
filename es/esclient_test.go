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
}

func (s *TestESSuit) SetupTest() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://124.222.48.125:9200/"))
	s.NoError(err)
	s.client = client

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

func TestES(t *testing.T) {
	suite.Run(t, new(TestESSuit))
}
