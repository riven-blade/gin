package domain

import "github.com/ddh-open/gin/app/module/base"

// DevopsSysDomain 域表
type DevopsSysDomain struct {
	base.DevopsModel
	Name     string `gorm:"column:name;type:varchar(256);unique;not null" json:"name"`           // 域账户名
	Enable   bool   `gorm:"column:enable;type:tinyint(1);default:null;default:1" json:"enable"`  // 是否启用
	UpdateID int64  `gorm:"column:update_id;type:bigint;default:null;default:0" json:"updateId"` // 更新人
	CreateID int64  `gorm:"column:create_id;type:bigint;default:null;default:0" json:"createId"` // 创建者
}
