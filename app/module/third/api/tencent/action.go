package tencent

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/third"
	"github.com/ddh-open/gin/framework/gin"
)

// TencentListResource godoc
// @Summary 获得腾讯云资源列表接口
// @Security ApiKeyAuth
// @Description 获得腾讯云资源列表接口
// @accept application/json
// @Produce application/json
// @Param data body third.TencentResourceListRequest true "页数，页大小，筛选条件"
// @Tags ThirdTencent
// @Success 200 {object}  base.Response
// @Router /third/tencent/resource/list [post]
func (t *ApiTencent) TencentListResource(c *gin.Context) {
	var param third.TencentResourceListRequest
	err := c.ShouldBindJSON(&param)
	res := base.Response{Code: 1, Msg: "查询成功", Data: nil}
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	result, err := t.service.GetTencentResourceList(param, c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc))
	if err != nil {
		res.Msg = err.Error()
		c.DJson(res)
		return
	}
	res.Data = result
	c.DJson(res)
}
