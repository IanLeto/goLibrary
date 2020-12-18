package gormDemo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB = NewDB()

type Client interface {
}

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3308)/gormloc2?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

// 创建表
func CreateTable(data interface{}) {
	DB.HasTable(data)
}
