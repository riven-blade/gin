package app

import (
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
}
