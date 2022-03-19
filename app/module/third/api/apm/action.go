package apm

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/third"
	"github.com/ddh-open/gin/framework/gin"
)

// AddMerchantApm godoc
// @Summary 新增商户apm
// @Security ApiKeyAuth
// @Description 新增商户apm
// @accept application/json
// @Produce application/json
// @Param data body third.AddMerchantApmRequest true "商户名称，商户id，名称空间"
// @Tags ThirdTencent
// @Success 200 {object}  base.Response
// @Router /third/tencent/apm/addMerchantApm [post]
func (t *ApiApm) AddMerchantApm(c *gin.Context) {
	var param third.AddMerchantApmRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "创建apm成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := t.service.AddMerchantApm(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result.Data
	c.DJson(res)
}
