package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goLibrary/utils"
)

var DB = NewMongoClient("localhost:27017", "ian")

func NewMongoClient(address, db string) *mgo.Database {
	url := fmt.Sprintf("mongodb://%s/%s", address, db)
	if url == "" {
		url = "mongodb://root:root.toor@172.18.38.86:27017/uhost_admin"
	}
	session, err := mgo.Dial(url)
	utils.NoErr(err)
	info, err := session.BuildInfo()
	if err != nil {
		logrus.Errorf("err %s", err)
	}
	fmt.Println(info)
	utils.NoErr(session.Ping())
	return session.DB("")
}
