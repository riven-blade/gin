package path

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/framework/gin"
)

// GetApi godoc
// @Summary 获得Api接口
// @Security ApiKeyAuth
// @Description 获得Api接口
// @accept application/json
// @Produce application/json
// @Param id path int true "Api id"
// @Tags Api
// @Success 200 {object}  base.Response
// @Router /sys/api/{id} [get]
func (a *ApiPath) GetApi(c *gin.Context) {
	roleId := c.Param("id")
	result, err := a.service.GetApiById(roleId, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ListApi godoc
// @Summary 获得Api列表接口
// @Security ApiKeyAuth
// @Description 获得Api列表接口
// @accept application/json
// @Produce application/json
// @Param data body base.PageRequest true "页数，页大小，筛选条件"
// @Tags Api
// @Success 200 {object}  base.Response
// @Router /sys/api/list [post]
func (a *ApiPath) ListApi(c *gin.Context) {
	var param base.PageRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "查询成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := a.service.GetApiList(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}

// AddApi godoc
// @Summary 新增Api接口
// @Security ApiKeyAuth
// @Description 新增Api接口
// @accept application/json
// @Produce application/json
// @Param data body path.DevopsSysApi true "Api"
// @Tags Api
// @Success 200 {object}  base.Response
// @Router /sys/api/add [post]
func (a *ApiPath) AddApi(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "新增成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.AddApi(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ModifyApi godoc
// @Summary 修改Api接口
// @Security ApiKeyAuth
// @Description 修改Api接口
// @accept application/json
// @Produce application/json
// @Param data body path.DevopsSysApi true "Api"
// @Tags Api
// @Success 200 {object}  base.Response
// @Router /sys/api/modify [post]
func (a *ApiPath) ModifyApi(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "修改成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.ModifyApi(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// DeleteApi godoc
// @Summary 删除Api接口
// @Security ApiKeyAuth
// @Description 删除Api接口
// @accept application/json
// @Produce application/json
// @Param ids body string true "Api ids"
// @Tags Api
// @Success 200 {object}  base.Response
// @Router /sys/api/delete [delete]
func (a *ApiPath) DeleteApi(c *gin.Context) {
	var ids string
	err := c.ShouldBindJSON(&ids)
	res := base.Response{Code: 1, Msg: "删除成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.DeleteApi(ids, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}
