package mongo_test

import (
	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestMongoSuit struct {
	suite.Suite
	session *mgo.Session
}

func (s *TestMongoSuit) SetupTest() {
	//s.session = mongo.NewMongoClient("mongodb://localhost:27017/uhost_admin")
}

func (s *TestMongoSuit) TestSimpleTest() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestMongoSuit))
}
