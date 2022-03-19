package group

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/sys/service/group"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiGroup struct {
	service *group.Service
}

func Register(r *gin.Engine) error {
	api := NewGroupApi(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	sysGroup := r.Group("/sys/", func(c *gin.Context) {
	})

	// 用户组相关接口
	sysGroup.GET("group/:id", api.GetGroups)
	sysGroup.GET("group/resource/:name", api.GetGroupsResource)
	sysGroup.POST("group/list", api.ListGroups)
	sysGroup.POST("group/add", api.AddGroup)
	sysGroup.POST("group/add/resources", api.AddResourcesToGroup)
	sysGroup.POST("group/modify", api.ModifyGroup)
	sysGroup.DELETE("group/delete", api.DeleteGroup)

	return nil
}

func NewGroupApi(c framework.Container) *ApiGroup {
	return &ApiGroup{service: group.NewService(c)}
}
