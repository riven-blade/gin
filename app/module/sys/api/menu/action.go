package menu

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/framework/gin"
)

// GetMenu godoc
// @Summary 获得菜单接口
// @Security ApiKeyAuth
// @Description 获得菜单接口
// @accept application/json
// @Produce application/json
// @Param id path int true "菜单id"
// @Tags Menu
// @Success 200 {object}  base.Response
// @Router /sys/menu/{id} [get]
func (a *ApiMenu) GetMenu(c *gin.Context) {
	id := c.Param("id")
	result, err := a.service.GetMenuById(id, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ListMenu godoc
// @Summary 获得菜单列表接口
// @Security ApiKeyAuth
// @Description 获得菜单列表接口
// @accept application/json
// @Produce application/json
// @Param data body base.PageRequest true "页数，页大小，筛选条件"
// @Tags Menu
// @Success 200 {object}  base.Response
// @Router /sys/menu/list [post]
func (a *ApiMenu) ListMenu(c *gin.Context) {
	var param base.PageRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "查询成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := a.service.GetMenuList(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}

// AddMenu godoc
// @Summary 新增菜单接口
// @Security ApiKeyAuth
// @Description 新增菜单接口
// @accept application/json
// @Produce application/json
// @Param data body menu.DevopsSysMenu true "菜单"
// @Tags Menu
// @Success 200 {object}  base.Response
// @Router /sys/menu/add [post]
func (a *ApiMenu) AddMenu(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "新增成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.AddMenu(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ModifyMenu godoc
// @Summary 修改菜单接口
// @Security ApiKeyAuth
// @Description 修改菜单接口
// @accept application/json
// @Produce application/json
// @Param data body menu.DevopsSysMenu true "菜单"
// @Tags Menu
// @Success 200 {object}  base.Response
// @Router /sys/menu/modify [post]
func (a *ApiMenu) ModifyMenu(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "修改成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.ModifyMenu(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// DeleteMenu godoc
// @Summary 删除菜单接口
// @Security ApiKeyAuth
// @Description 删除菜单接口
// @accept application/json
// @Produce application/json
// @Param ids body string true "菜单ids"
// @Tags Menu
// @Success 200 {object}  base.Response
// @Router /sys/menu/delete [delete]
func (a *ApiMenu) DeleteMenu(c *gin.Context) {
	var ids string
	err := c.ShouldBindJSON(&ids)
	res := base.Response{Code: 1, Msg: "删除成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.DeleteMenu(ids, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}
