package gormDemo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB = NewDB()

type Client interface {
}

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/ianDB?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

// 创建表
func CreateTable(data interface{}) {
	DB.HasTable(data)
}

// 定义orm  注意字段要大写
type user struct {
	Id       int    `gorm:"id"`
	Name     string `gorm:"name"`
	Follower string `gorm:"follower"`
}

// 常用语法
// 条件查询
func SelectByID() []user {
	var (
		err error
		// 定义输入/输出
		user []user
	)
	// 查询全部
	err = DB.Table("user").Find(&user).Error
	if err != nil {
		panic(err)
	}
	// 格式化结果
	return user
}
