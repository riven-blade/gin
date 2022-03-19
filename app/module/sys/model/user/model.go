package user

import (
	"github.com/ddh-open/gin/app/module/base"
	uuid "github.com/satori/go.uuid"
)

// SysUser 用户
type SysUser struct {
	base.DevopsModel
	UUID         uuid.UUID `gorm:"column:uuid;type:varchar(32);default:null"`                  // UUID
	Username     string    `gorm:"unique;column:username;type:varchar(32);not null"`           // 登录名/11111
	Password     string    `gorm:"column:password;type:varchar(32);not null"`                  // 密码
	Salt         string    `gorm:"column:salt;type:varchar(16);not null;default:1111"`         // 密码盐
	RealName     string    `gorm:"column:real_name;type:varchar(32);default:null"`             // 真实姓名
	DepartID     int       `gorm:"column:depart_id;type:int;default:null;default:0"`           // 部门/11111/dict
	UserType     int       `gorm:"column:user_type;type:int;default:null;default:2"`           // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status       int       `gorm:"column:status;type:int;default:null;default:10"`             // 状态
	Thirdid      string    `gorm:"column:thirdid;type:varchar(200);default:null"`              // 第三方ID
	Endtime      string    `gorm:"column:endtime;type:varchar(32);default:null"`               // 结束时间
	Email        string    `gorm:"column:email;type:varchar(64);default:null"`                 // email
	Tel          string    `gorm:"column:tel;type:varchar(32);default:null"`                   // 手机号
	Address      string    `gorm:"column:address;type:varchar(32);default:null"`               // 地址
	TitleURL     string    `gorm:"column:title_url;type:varchar(200);default:null"`            // 头像地址
	Remark       string    `gorm:"column:remark;type:varchar(1000);default:null"`              // 说明
	Theme        string    `gorm:"column:theme;type:varchar(64);default:null;default:default"` // 主题
	BackSiteID   int       `gorm:"column:back_site_id;type:int;default:null;default:0"`        // 后台选择站点ID
	CreateSiteID int       `gorm:"column:create_site_id;type:int;default:null;default:1"`      // 创建站点ID
	ProjectID    int64     `gorm:"column:project_id;type:bigint;default:null;default:0"`       // 项目ID
	ProjectName  string    `gorm:"column:project_name;type:varchar(100);default:null"`         // 项目名称
	Enable       bool      `gorm:"column:enable;type:tinyint(1);default:null;default:1"`       // 是否启用//radio/1,启用,2,禁用
	UpdateID     int64     `gorm:"column:update_id;type:bigint;default:null;default:0"`        // 更新人
	CreateID     int64     `gorm:"column:create_id;type:bigint;default:null;default:0"`        // 创建者
}

// SysUserColumns get sql column name.获取数据库列名
var SysUserColumns = struct {
	ID           string
	UUID         string
	Username     string
	Password     string
	Salt         string
	RealName     string
	DepartID     string
	UserType     string
	Status       string
	Thirdid      string
	Endtime      string
	Email        string
	Tel          string
	Address      string
	TitleURL     string
	Remark       string
	Theme        string
	BackSiteID   string
	CreateSiteID string
	ProjectID    string
	ProjectName  string
	Enable       string
	UpdateTime   string
	UpdateID     string
	CreateTime   string
	CreateID     string
	DeleteTime   string
}{
	ID:           "id",
	UUID:         "uuid",
	Username:     "username",
	Password:     "password",
	Salt:         "salt",
	RealName:     "real_name",
	DepartID:     "depart_id",
	UserType:     "user_type",
	Status:       "status",
	Thirdid:      "thirdid",
	Endtime:      "endtime",
	Email:        "email",
	Tel:          "tel",
	Address:      "address",
	TitleURL:     "title_url",
	Remark:       "remark",
	Theme:        "theme",
	BackSiteID:   "back_site_id",
	CreateSiteID: "create_site_id",
	ProjectID:    "project_id",
	ProjectName:  "project_name",
	Enable:       "enable",
	UpdateTime:   "update_time",
	UpdateID:     "update_id",
	CreateTime:   "create_time",
	CreateID:     "create_id",
	DeleteTime:   "delete_time",
}
