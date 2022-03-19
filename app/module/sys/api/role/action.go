package role

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/framework/gin"
)

// GetRolesResource godoc
// @Summary 获得角色权限接口
// @Security ApiKeyAuth
// @Description 获得角色权限接口
// @accept application/json
// @Produce application/json
// @Param name path string true "角色name"
// @Param domain query string false "域"
// @Tags Group
// @Success 200 {object}  base.Response
// @Router /sys/group/role/{name} [get]
func (a *ApiRole) GetRolesResource(c *gin.Context) {
	name := c.Param("name")
	domain := c.Query("domain")
	// 从cookie中获取domain
	if d, err := c.Cookie("domain"); err == nil {
		domain = d
	}
	result, err := a.service.GetRolesResource(name, domain, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// GetRoles godoc
// @Summary 获得角色接口
// @Security ApiKeyAuth
// @Description 获得角色接口
// @accept application/json
// @Produce application/json
// @Param id path int true "角色id"
// @Tags Role
// @Success 200 {object}  base.Response
// @Router /sys/roles/{id} [get]
func (a *ApiRole) GetRoles(c *gin.Context) {
	roleId := c.Param("id")
	result, err := a.service.GetRoleById(roleId, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	res := base.Response{Code: 1, Msg: "查询成功", Data: result}
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ListRoles godoc
// @Summary 获得角色列表接口
// @Security ApiKeyAuth
// @Description 获得角色列表接口
// @accept application/json
// @Produce application/json
// @Param data body base.PageRequest true "页数，页大小，筛选条件"
// @Tags Role
// @Success 200 {object}  base.Response
// @Router /sys/roles/list [post]
func (a *ApiRole) ListRoles(c *gin.Context) {
	var param base.PageRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "查询成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := a.service.GetRoleList(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}

// AddRole godoc
// @Summary 新增角色接口
// @Security ApiKeyAuth
// @Description 新增角色接口
// @accept application/json
// @Produce application/json
// @Param data body role.DevopsSysRole true "角色"
// @Tags Role
// @Success 200 {object}  base.Response
// @Router /sys/roles/add [post]
func (a *ApiRole) AddRole(c *gin.Context) {
	param := make(map[string]interface{}, 0)
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "新增成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.AddRole(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// ModifyRole godoc
// @Summary 修改角色接口
// @Security ApiKeyAuth
// @Description 修改角色接口
// @accept application/json
// @Produce application/json
// @Param data body role.DevopsSysRole true "角色"
// @Tags Role
// @Success 200 {object}  base.Response
// @Router /sys/roles/modify [post]
func (a *ApiRole) ModifyRole(c *gin.Context) {
	var param map[string]interface{}
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "修改成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.ModifyRole(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// DeleteRole godoc
// @Summary 删除角色接口
// @Security ApiKeyAuth
// @Description 删除角色接口
// @accept application/json
// @Produce application/json
// @Param ids body string true "角色ids"
// @Tags Role
// @Success 200 {object}  base.Response
// @Router /sys/roles/delete [delete]
func (a *ApiRole) DeleteRole(c *gin.Context) {
	var ids string
	err := c.ShouldBindJSON(&ids)
	res := base.Response{Code: 1, Msg: "删除成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.DeleteRole(ids, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}

// AddResourcesToRole godoc
// @Summary 给角色新增权限接口
// @Security ApiKeyAuth
// @Description 给角色新增权限接口
// @accept application/json
// @Produce application/json
// @Param data body []base.CabinInReceive true "新增api权限Ptype为p； 新增菜单权限Ptype为p3 , source 是角色的id，resource 是资源"
// @Tags Role
// @Success 200 {object}  base.Response
// @Router /sys/roles/add/resources [post]
func (a *ApiRole) AddResourcesToRole(c *gin.Context) {
	param := make([]base.CabinInReceive, 0)
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "新增成功"}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	err = a.service.AddResourcesToRole(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
	}
	c.DJson(res)
}
