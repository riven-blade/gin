package user

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/sys"
	"github.com/ddh-open/gin/framework/gin"
)

// UserRelativeGroup godoc
// @Summary 获得用户分组资源接口
// @Security ApiKeyAuth
// @Description 获得用户分组资源接口
// @accept application/json
// @Produce application/json
// @Param id path int true "用户id"
// @Param domain query string false "域"
// @Tags User
// @Success 200 {object}  base.Response
// @Router /user/relative/group/{id} [get]
func (api *ApiUser) UserRelativeGroup(c *gin.Context) {
	id := c.Param("id")
	domain := c.Query("domain")
	// 从cookie中获取domain
	if d, err := c.Cookie("domain"); err == nil {
		domain = d
	}
	result, err := api.service.GetGroupsByUserId(id, domain, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// UserRelativeGroupAdd godoc
// @Summary 新增用户分组接口
// @Security ApiKeyAuth
// @Description 新增用户分组接口
// @accept application/json
// @Produce application/json
// @Param data body sys.RelativeUserRequest true "关联参数"
// @Tags User
// @Success 200 {object}  base.Response
// @Router /user/relative/group/add [post]
func (api *ApiUser) UserRelativeGroupAdd(c *gin.Context) {
	param := sys.RelativeUserRequest{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "新增成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	// 从cookie中获取domain
	if domain, err := c.Cookie("domain"); err == nil {
		param.Domain = domain
	}
	if param.Domain == "" {
		param.Domain = "1"
	}
	err = api.service.RelativeGroupsToUser(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// UserRelativeGroupDelete godoc
// @Summary 删除用户分组接口
// @Security ApiKeyAuth
// @Description 删除用户分组接口
// @accept application/json
// @Produce application/json
// @Param data body sys.RelativeUserRequest true "关联参数"
// @Tags User
// @Success 200 {object}  base.Response
// @Router /user/relative/group/delete [delete]
func (api *ApiUser) UserRelativeGroupDelete(c *gin.Context) {
	param := sys.RelativeUserRequest{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "删除成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	// 从cookie中获取domain
	if domain, err := c.Cookie("domain"); err == nil {
		param.Domain = domain
	}
	if param.Domain == "" {
		param.Domain = "1"
	}
	err = api.service.DeleteRelativeGroupsToUser(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}
