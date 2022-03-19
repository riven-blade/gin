package path

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/sys/service/path"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiPath struct {
	service *path.Service
}

func Register(r *gin.Engine) error {
	api := NewSysApi(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	sysGroup := r.Group("/sys/", func(c *gin.Context) {
	})

	// api相关接口
	sysGroup.GET("api/:id", api.GetApi)
	sysGroup.POST("api/list", api.ListApi)
	sysGroup.POST("api/add", api.AddApi)
	sysGroup.POST("api/modify", api.ModifyApi)
	sysGroup.DELETE("api/delete", api.DeleteApi)

	return nil
}

func NewSysApi(c framework.Container) *ApiPath {
	return &ApiPath{service: path.NewService(c)}
}
