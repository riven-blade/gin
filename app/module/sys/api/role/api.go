package role

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/sys/service/role"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiRole struct {
	service *role.Service
}

func Register(r *gin.Engine) error {
	api := NewSysApi(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	sysGroup := r.Group("/sys/", func(c *gin.Context) {
	})

	// 用户角色相关接口
	sysGroup.GET("roles/:id", api.GetRoles)
	sysGroup.GET("roles/resource/:name", api.GetRolesResource)
	sysGroup.POST("roles/list", api.ListRoles)
	sysGroup.POST("roles/add", api.AddRole)
	sysGroup.POST("roles/add/resources", api.AddResourcesToRole)
	sysGroup.POST("roles/modify", api.ModifyRole)
	sysGroup.DELETE("roles/delete", api.DeleteRole)

	return nil
}

func NewSysApi(c framework.Container) *ApiRole {
	return &ApiRole{service: role.NewService(c)}
}
