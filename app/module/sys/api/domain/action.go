package domain

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/framework/gin"
)

// GetDomains godoc
// @Summary 获得域接口
// @Security ApiKeyAuth
// @Description 获得域接口
// @accept application/json
// @Produce application/json
// @Param id path int true "域id"
// @Tags Domain
// @Success 200 {object}  base.Response
// @Router /sys/domain/{id} [get]
func (a *ApiDomain) GetDomains(c *gin.Context) {
	id := c.Param("id")
	result, err := a.service.GetDomainById(id, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ListDomains godoc
// @Summary 获得域列表接口
// @Security ApiKeyAuth
// @Description 获得域列表接口
// @accept application/json
// @Produce application/json
// @Param data body base.PageRequest true "页数，页大小，筛选条件"
// @Tags Domain
// @Success 200 {object}  base.Response
// @Router /sys/domain/list [post]
func (a *ApiDomain) ListDomains(c *gin.Context) {
	var param base.PageRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "查询成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := a.service.GetDomainList(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}

// AddDomain godoc
// @Summary 新增域接口
// @Security ApiKeyAuth
// @Description 新增域接口
// @accept application/json
// @Produce application/json
// @Param data body domain.DevopsSysDomain true "域"
// @Tags Domain
// @Success 200 {object}  base.Response
// @Router /sys/domain/add [post]
func (a *ApiDomain) AddDomain(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "新增成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.AddDomain(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ModifyDomain godoc
// @Summary 修改域接口
// @Security ApiKeyAuth
// @Description 修改域接口
// @accept application/json
// @Produce application/json
// @Param data body DevopsSysDomain true "域"
// @Tags Domain
// @Success 200 {object}  base.Response
// @Router /sys/domain/modify [post]
func (a *ApiDomain) ModifyDomain(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "修改成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.ModifyDomain(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// DeleteDomain godoc
// @Summary 删除域接口
// @Security ApiKeyAuth
// @Description 删除域接口
// @accept application/json
// @Produce application/json
// @Param ids body string true "域ids"
// @Tags Domain
// @Success 200 {object}  base.Response
// @Router /sys/domain/delete [delete]
func (a *ApiDomain) DeleteDomain(c *gin.Context) {
	var ids string
	err := c.ShouldBindJSON(&ids)
	res := base.Response{Code: 1, Msg: "删除成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.DeleteDomain(ids, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}
