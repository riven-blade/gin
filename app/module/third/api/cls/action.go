package cls

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/third"
	"github.com/ddh-open/gin/framework/gin"
)

// AddMerchantClsLogTopic godoc
// @Summary 新增商户的日志主题和日志集
// @Security ApiKeyAuth
// @Description 新增商户的日志主题和日志集
// @accept application/json
// @Produce application/json
// @Param data body third.AddMerchantClsLogTopicRequest true "商户名称，商户id，名称空间"
// @Tags ThirdTencent
// @Success 200 {object}  base.Response
// @Router /third/tencent/cls/topic/addMerchantClsLogTopic [post]
func (t *ApiCls) AddMerchantClsLogTopic(c *gin.Context) {
	var param third.AddMerchantClsLogTopicRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "创建成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := t.service.AddMerchantClsLogTopic(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}

// DeleteMerchantLog godoc
// @Summary 删除商户日志主题和日志集
// @Security ApiKeyAuth
// @Description 删除商户日志主题和日志集
// @accept application/json
// @Produce application/json
// @Param data body third.DeleteMerchantLog true "商户名称，商户id"
// @Tags ThirdTencent
// @Success 200 {object}  base.Response
// @Router /third/tencent/cls/topic/deleteMerchantLog [post]
func (t *ApiCls) DeleteMerchantLog(c *gin.Context) {
	var param third.DeleteMerchantLog
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "删除成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := t.service.DeleteMerchantLog(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}
