package menu

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/sys/service/menu"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiMenu struct {
	service *menu.Service
}

func Register(r *gin.Engine) error {
	api := NewSysApi(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	sysGroup := r.Group("/sys/", func(c *gin.Context) {
	})

	// 菜单相关接口
	sysGroup.GET("menu/:id", api.GetMenu)
	sysGroup.POST("menu/list", api.ListMenu)
	sysGroup.POST("menu/add", api.AddMenu)
	sysGroup.POST("menu/modify", api.ModifyMenu)
	sysGroup.DELETE("menu/delete", api.DeleteMenu)

	return nil
}

func NewSysApi(c framework.Container) *ApiMenu {
	return &ApiMenu{service: menu.NewService(c)}
}
