package path

import "github.com/ddh-open/gin/app/module/base"

// DevopsSysApi Api 表
type DevopsSysApi struct {
	base.DevopsModel
	Path        string `json:"path" gorm:"comment:api路径;unique;not null"`                           // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`                                  // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`                                        // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"`                               // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	UpdateID    int64  `gorm:"column:update_id;type:bigint;default:null;default:0" json:"updateId"` // 更新人
	CreateID    int64  `gorm:"column:create_id;type:bigint;default:null;default:0" json:"createId"` // 创建者
}
