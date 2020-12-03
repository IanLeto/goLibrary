package gormDemo_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"goLibrary/gormDemo"
	"testing"
)

type TestDBSuit struct {
	suite.Suite
	db *gorm.DB
}

func (s *TestDBSuit) SetupTest() {
	s.db = gormDemo.NewDB()
}

//
func (s *TestDBSuit) TestCreatFunc() {
	s.db.CreateTable(&gormDemo.User{})
	s.Equal(true, s.db.HasTable(&gormDemo.User{}))
}

// 关联演示
type UserDemo struct {
	gorm.Model
	Profile      Profile `gorm:"ForeignKey:ProfileRefer"`
	ProfileRefer int
	CreditCard   CreditCardDemo
	EmailID      uint
	Emails       []EmailDemo `gorm:"ForeignKey:EmailID;AssociationForeignKey:EmailID"`
}

type Profile struct {
	gorm.Model
	Name string
}

func (s *TestDBSuit) TestKeysFunc() {
	//s.db.CreateTable(&UserDemo{})
	//s.db.CreateTable(&Profile{})

}

// 包含一个
type CreditCardDemo struct {
	gorm.Model
	UserID uint
}

// 修改表名
func (s CreditCardDemo) TableName() string {
	return "creditCardDemo"
}

// 修改表结构
// 增加列（增加字段）直接在结构提里面继上一个字段就成

func (s *TestDBSuit) TestAlterTable() {
}

func (s *TestDBSuit) TestCreditCardNewRecordFunc() {
	if !s.db.HasTable(&CreditCardDemo{}) {
		s.db.CreateTable(&CreditCardDemo{})
	}
	// 新建一条数据
	s.db.Create(&CreditCardDemo{
		Model:  gorm.Model{},
		UserID: 1,
	})
	// 新建带关联而不是外键的数据
	s.db.Create(&UserDemo{
		Model:        gorm.Model{},
		Profile:      Profile{},
		ProfileRefer: 0,
		CreditCard: CreditCardDemo{
			Model:  gorm.Model{},
			UserID: 12,
		},
	})
	// fetch 数据
	var user UserDemo

	s.db.Model(&user)
	s.Equal(uint(0), user.ID)
	//  fetch 关联 不指定关联，查询毫无意义？？？
	var card CreditCardDemo
	s.db.Model(&user).Related(&card)
	//s.Equal(uint(1), card.ID)

}

// one2many
type EmailDemo struct {
	gorm.Model
	Email   string
	UserID  uint
	Name    string
	EmailID uint
}

// 创建带外键的记录
func (s *TestDBSuit) TestManySuit() {
	if !s.db.HasTable("email_demos") {
		s.db.CreateTable(&EmailDemo{})
	}
	// migrate db
	s.db.AutoMigrate(&UserDemo{})
	s.db.AutoMigrate(&EmailDemo{})
	var user = UserDemo{
		EmailID: uint(10),
		Emails: []EmailDemo{
			{
				Email:   "11",
				UserID:  uint(1),
				EmailID: uint(10),
			},
			{
				Email:   "21",
				UserID:  uint(2),
				EmailID: uint(10),
			},
		},
	}
	s.NoError(s.db.Create(&user).Error)
}

// 关联查询 常用！！！
// 单独外键
type Main struct {
	gorm.Model
	Name  string
	KeyID string
	Key   Foreigner `gorm:"ForeignKey:KeyID;AssociationForeignKey:Room"`
}

type Foreigner struct {
	Name string
	Room string
	gorm.Model
}

func (s TestDBSuit) TestCreatTableTest() {
	if !s.db.HasTable(&Main{}) {
		s.db.CreateTable(&Main{})
		s.db.CreateTable(&Foreigner{})
	}
}

// 带有外键的存储
func (s *TestDBSuit) TestAssociationCreateTest() {

	s.db.AutoMigrate(&Main{})
	s.db.AutoMigrate(&Foreigner{})
	s.db.Create(&Main{
		Name: "sicong",
		// 不用+ 上keyID 这个字段
		Key: Foreigner{
			Name: "littleSan",
			Room: "k003",
		},
	})
}

// 带有外键的查询
func (s *TestDBSuit) TestAssociationQuery() {
	var (
		userWang Main
		//little3   Foreigner
	)
	s.db.Model(&userWang).Related(&userWang.Key, "Room").Where("key_id = 3")
	fmt.Println(userWang.Key.Name)
}

// one2many
type Leader struct {
	LeaderID           int    `gorm:"primary_key"`
	Name               string `gorm:"name"`
	ForeignKeyLeaderID int
	Followers          []Follower `gorm:"ForeignKey:ForeignKeyLeaderID"`
}
type Follower struct {
	FollowerID         int    `gorm:"primary_key"`
	Name               string `gorm:"name"`
	ForeignKeyLeaderID int
}

func (s TestDBSuit) TestOne2many() {
	if !s.db.HasTable(&Leader{}) {
		s.db.CreateTable(&Leader{})
		s.db.CreateTable(&Follower{})
	}
	s.db.AutoMigrate(&Leader{}, &Follower{})
	s.db.Create(&Leader{
		Name: "leader23",
		Followers: []Follower{
			{
				FollowerID: 122,
				Name:       "x12",
			}, {
				FollowerID: 123,
				Name:       "x33",
			},
		},
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestDBSuit))
}
