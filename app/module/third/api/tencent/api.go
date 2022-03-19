package tencent

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/third/service/tencent"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiTencent struct {
	service *tencent.Service
}

func Register(r *gin.Engine) error {
	api := NewThirdApi(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	userGroup := r.Group("/third/", func(c *gin.Context) {
	})
	// 腾讯云相关接口
	userGroup.POST("tencent/resource/list", api.TencentListResource) // 获取腾讯云的资源（主机，数据库，redis and soon）
	return nil
}

func NewThirdApi(c framework.Container) *ApiTencent {
	return &ApiTencent{service: tencent.NewService(c)}
}
