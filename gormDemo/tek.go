package gormDemo

import (
	"database/sql"
	"time"
)

type User struct {
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增

	CreditCard CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails     []Email    // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress   Address // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	UserID uint
	Number string
}

// Template 表
type Template struct {
	Name              string          `gorm:"type:varchar(64);not null;" json:"Name" binding:"required"`             // Template 名
	TemplateID        string          `gorm:"type:varchar(64);not null" json:"TemplateID"`                           // Template 的uuid
	ProjectName       string          `gorm:"type:varchar(64);not null;" json:"ProjectName" binding:"required"`      // 项目名
	DeploymentService string          `gorm:"type:varchar(64);not null" json:"DeploymentService" binding:"required"` // 部署 agent 服务 URL 的uuid
	ServiceName       string          `gorm:"type:varchar(64);not null" json:"ServiceName" binding:"required"`       // playbook yaml
	Owner             string          `gorm:"type:varchar(16)" json:"Owner" `                                        // 创建者
	Comm              string          `gorm:"type:varchar(64)" json:"Comm"`                                          // 备注
	CheckList         []TemplateCheck `gorm:"ForeignKey:TemplateID;AssociationForeignKey:TemplateID"`
}

// TemplateCheck 检查表
type TemplateCheck struct {
	TemplateID string `gorm:"type:varchar(64);not null" json:"TemplateID" binding:"required"` // Template 的uuid
	CheckName  string `gorm:"type:varchar(64);not null" json:"CheckName" binding:"required"`  // 检查 Name
	CheckID    string `gorm:"type:varchar(64);not null" json:"CheckID" binding:"required"`    // 检查 ID
	Type       string `gorm:"type:varchar(64);not null" json:"Type" binding:"required"`       // 检查类型  checkpoint,gamma, drill, custom
	ProcessID  string `gorm:"type:varchar(64);not null" json:"ProcessID" binding:"required"`
}

// Task ...
type Task struct {
	Name          string    `gorm:"type:varchar(64);not null;" json:"Name" binding:"required"`       //Task 名
	TaskID        string    `gorm:"type:varchar(64);not null;primary_key" json:"TaskID"`             //Task 的uuid
	TemplateID    string    `gorm:"type:varchar(64);not null;" json:"TemplateID" binding:"required"` //Template 的uuid
	SubTaskList   []SubTask `gorm:"ForeignKey:TaskID;AssociationForeignKey:TaskID"`
	Status        int       `gorm:"type:int(16);not null;default:0" json:"Status"`        //  0: 未开始, 1 :进行中, 2 :成功, 3 :失败, 4 :异常
	JobID         string    `gorm:"type:varchar(64);;" json:"JobID"`                      // agent job 的 id
	Owner         string    `gorm:"type:varchar(16)" json:"Owner" `                       // 创建者
	Comm          string    `gorm:"type:varchar(64)" json:"Comm"`                         //备注
	SecurityCheck int       `gorm:"type:int(16);not null;default:0" json:"SecurityCheck"` //  0: 未开始, 1 :进行中, 2 :成功, 3 :失败, 4 :异常
	Acceptance    int       `gorm:"type:int(16);not null;default:0" json:"Acceptance"`    //  0: 未开始, 1 :进行中, 2 :成功, 3 :失败, 4 :异常
	ServiceCheck  int       `gorm:"type:int(16);not null;default:0" json:"ServiceCheck"`  //  0: 未开始, 1 :进行中, 2 :成功, 3 :失败, 4 :异常
	Test          int       `gorm:"type:int(16);not null;default:0" json:"Test"`          //  0: 未开始, 1 :进行中, 2 :成功, 3 :失败, 4 :异常
}

// SubTask 表
type SubTask struct {
	SubTaskID string         `gorm:"type:varchar(64);not null;" json:"SubTaskID"`   //SubTask 的uuid
	IP        string         `gorm:"type:varchar(16)" json:"IP" `                   //子任务执行的IP
	TaskID    string         `gorm:"type:varchar(64);not null;" json:"TaskID"`      //Task 的uuid
	Status    int            `gorm:"type:int(16);not null;default:0" json:"Status"` //  0: 未开始, 1 :进行中, 2 :成功, 3 :失败, 4 :异常
	ParamList []SubTaskParam `gorm:"ForeignKey:SubTaskID;AssociationForeignKey:SubTaskID"`
}

type SubTaskParam struct {
	SubTaskID  string `gorm:"type:varchar(64);not null;" json:"SubTaskID"`                   //Task 的uuid
	IP         string `gorm:"type:varchar(16)" json:"IP" binding:"required"`                 //使用此param的ip
	ParamKey   string `gorm:"type:varchar(64);not null;" json:"ParamKey" binding:"required"` //参数名, 模板字面量 使用${}
	ParamValue string `gorm:"type:text" json:"ParamValue" binding:"required"`                //参数值,模板字面量 使用${}
	Comm       string `gorm:"type:varchar(255)" json:"Comm"`                                 //备注
}
