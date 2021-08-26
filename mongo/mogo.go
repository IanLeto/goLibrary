package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goLibrary/utils"
)

func NewMongoClient(url string) *mgo.Session {
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
	return session
}

// bson 演示
func FindOne() {

}
