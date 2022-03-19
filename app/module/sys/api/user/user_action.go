package user

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/sys"
	"github.com/ddh-open/gin/framework/gin"
)

// Login godoc
// @Summary 用户登录接口
// @Security ApiKeyAuth
// @Description 用户登录接口
// @accept application/json
// @Produce application/json
// @Param data body sys.LoginRequest true "用户名，密码，账户类型"
// @Tags User
// @Success 200 {object}  base.Response
// @Router /user/login [post]
func (api *ApiUser) Login(c *gin.Context) {
	logger := c.MustMakeLog()
	var request sys.LoginRequest
	err := c.ShouldBindJSON(&request)
	res := base.Response{
		Code: 1,
		Msg:  "",
		Data: nil,
	}
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	grpc := c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc)
	data, err := api.service.Login(request, grpc)
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	res.Data = data
	c.DJson(res)
}

// Register godoc
// @Summary 用户注册接口
// @Security ApiKeyAuth
// @Description 用户注册接口
// @Produce  json
// @Tags User
// @Success 200 {object} base.Response
// @Router /user/register [post]
func (api *ApiUser) Register(c *gin.Context) {
	logger := c.MustMakeLog()
	var request sys.LoginRequest
	err := c.ShouldBindJSON(&request)
	res := base.Response{
		Code: 1,
		Msg:  "",
		Data: nil,
	}
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	grpc := c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc)
	data, err := api.service.Login(request, grpc)
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	res.Data = data
	c.DJson(res)
}

// Logout godoc
// @Summary 用户退出接口
// @Security ApiKeyAuth
// @Description 用户退出接口
// @Produce  json
// @Tags User
// @Success 200 {object} base.Response
// @Router /user/logout [post]
func (api *ApiUser) Logout(c *gin.Context) {
	logger := c.MustMakeLog()
	var request sys.LoginRequest
	err := c.ShouldBindJSON(&request)
	res := base.Response{
		Code: 1,
		Msg:  "",
		Data: nil,
	}
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	grpc := c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc)
	data, err := api.service.Login(request, grpc)
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	res.Data = data
	c.DJson(res)
}

// GetUserInfo godoc
// @Summary 获取用户详情的接口
// @Security ApiKeyAuth
// @Description 根据token获取id
// @Produce  json
// @Tags User
// @Success 200 {object} base.Response
// @Router /user/info [get]
func (api *ApiUser) GetUserInfo(c *gin.Context) {
	logger := c.MustMakeLog()
	data, err := api.service.GetUserInfo(c)
	res := base.Response{
		Code: 1,
		Data: data,
	}
	if err != nil {
		logger.Error(err.Error())
		res.Code = -1
		res.Msg = err.Error()
	}
	c.DJson(res)
}
