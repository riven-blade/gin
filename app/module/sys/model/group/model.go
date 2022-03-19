package group

import "github.com/ddh-open/gin/app/module/base"

// DevopsSysGroup 域表
type DevopsSysGroup struct {
	base.DevopsModel
	ParentID  int    `gorm:"column:parent_id;type:int;default:null;default:0" json:"parentId"`   // 上级机构
	Name      string `gorm:"unique;column:name;type:varchar(32);not null" json:"name"`           // 部门/11111
	Code      string `gorm:"column:code;type:varchar(128);default:null" json:"code"`             // 机构编码
	Sort      int    `gorm:"column:sort;type:int;default:null;default:0" json:"sort"`            // 序号
	Linkman   string `gorm:"column:linkman;type:varchar(64);default:null" json:"linkman"`        // 联系人
	LinkmanNo string `gorm:"column:linkman_no;type:varchar(32);default:null" json:"linkmanNo"`   // 联系人电话
	Remark    string `gorm:"column:remark;type:varchar(128);default:null" json:"remark"`         // 组描述
	Enable    bool   `gorm:"column:enable;type:tinyint(1);default:null;default:1" json:"enable"` // 是否启用
	Alias     string `gorm:"column:alias;type:varchar(128);default:null" json:"alias"`           // 别名
	Wechat    string `gorm:"column:wechat;type:varchar(128);default:null" json:"wechat"`         // wechat
	Domain    int    `json:"domain" gorm:"column:domain;type:int;default:null"`                  // 域
}
