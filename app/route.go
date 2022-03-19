package app

import (
	"github.com/ddh-open/gin/app/module/sys/api/domain"
	"github.com/ddh-open/gin/app/module/sys/api/group"
	"github.com/ddh-open/gin/app/module/sys/api/menu"
	"github.com/ddh-open/gin/app/module/sys/api/path"
	"github.com/ddh-open/gin/app/module/sys/api/role"
	"github.com/ddh-open/gin/app/module/sys/api/user"
	"github.com/ddh-open/gin/app/module/third/api/apm"
	"github.com/ddh-open/gin/app/module/third/api/cls"
	"github.com/ddh-open/gin/app/module/third/api/tencent"
	"github.com/ddh-open/gin/app/swagger"
	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/gin"
	ginSwagger "github.com/ddh-open/gin/framework/middleware/gin-swagger"
	"github.com/ddh-open/gin/framework/middleware/gin-swagger/swaggerFiles"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	// set swagger info
	swagger.SwaggerInfo.Title = "github.com/ddh-open/gin API"
	swagger.SwaggerInfo.Description = "This is github.com/ddh-open/gin API"
	swagger.SwaggerInfo.Version = "1.0"
	swagger.SwaggerInfo.Host = ""
	swagger.SwaggerInfo.BasePath = ""
	swagger.SwaggerInfo.Schemes = []string{"http"}
	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	/** 系统相关  **/
	domain.Register(r)
	group.Register(r)
	menu.Register(r)
	path.Register(r)
	role.Register(r)
	// 用户模块注册路由
	user.Register(r)

	/** 第三方相关  **/
	apm.Register(r)
	cls.Register(r)
	tencent.Register(r)
}
