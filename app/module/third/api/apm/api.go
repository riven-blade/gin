package apm

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/third/service/apm"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiApm struct {
	service *apm.Service
}

func Register(r *gin.Engine) error {
	api := NewApiApm(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	userGroup := r.Group("/third/", func(c *gin.Context) {
	})
	// apm相关接口
	userGroup.POST("tencent/apm/addMerchantApm", api.AddMerchantApm)
	return nil
}

func NewApiApm(c framework.Container) *ApiApm {
	return &ApiApm{service: apm.NewService(c)}
}
