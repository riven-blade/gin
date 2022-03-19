package user

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/framework/gin"
)

// UserRelativeMenu godoc
// @Summary 获得用户菜单接口
// @Security ApiKeyAuth
// @Description 获得用户菜单接口
// @accept application/json
// @Produce application/json
// @Param id path int true "用户id"
// @Param domain query string false "域"
// @Tags User
// @Success 200 {object}  base.Response
// @Router /user/relative/menu/{id} [get]
func (api *ApiUser) UserRelativeMenu(c *gin.Context) {
	uuid := c.Param("uuid")
	domain := c.Query("domain")
	// 从cookie中获取domain
	if d, err := c.Cookie("domain"); err == nil {
		domain = d
	}
	result, err := api.service.GetMenusByUserId(uuid, domain, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}
