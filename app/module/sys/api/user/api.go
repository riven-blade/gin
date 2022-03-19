package user

import (
	"github.com/ddh-open/gin/app/middleware/auth"
	"github.com/ddh-open/gin/app/module/sys/service/user"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/gin"
)

type ApiUser struct {
	service *user.Service
}

func Register(r *gin.Engine) error {
	api := NewUserApi(r.GetContainer())
	r.Use(auth.MiddlewareAuth())
	userGroup := r.Group("/user/", func(c *gin.Context) {
	})
	userGroup.POST("login", api.Login)       // 用户登录
	userGroup.POST("logout", api.Logout)     // 用户退出
	userGroup.POST("register", api.Register) // 用户注册

	// 用户关联角色相关
	userGroup.GET("relative/role/:id", api.UserRelativeRole)             // 获取用户的相关权限
	userGroup.POST("relative/role/add", api.UserRelativeRoleAdd)         // 给用户增加角色
	userGroup.DELETE("relative/role/delete", api.UserRelativeRoleDelete) // 给用户删除角色
	// 用户关联分组
	userGroup.GET("relative/group/:id", api.UserRelativeGroup)             // 获取用户的相关权限
	userGroup.POST("relative/group/add", api.UserRelativeGroupAdd)         // 给用户增加角色
	userGroup.DELETE("relative/group/delete", api.UserRelativeGroupDelete) // 给用户删除角色

	return nil
}

func NewUserApi(c framework.Container) *ApiUser {
	return &ApiUser{service: user.NewService(c)}
}
