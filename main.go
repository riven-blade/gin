package main

import (
	app2 "github.com/ddh-open/gin/app"
	"github.com/ddh-open/gin/app/provider/grpc"
	"github.com/ddh-open/gin/boot"
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/provider/app"
	"github.com/ddh-open/gin/framework/provider/cache"
	"github.com/ddh-open/gin/framework/provider/config"
	"github.com/ddh-open/gin/framework/provider/env"
	"github.com/ddh-open/gin/framework/provider/id"
	"github.com/ddh-open/gin/framework/provider/kernel"
	"github.com/ddh-open/gin/framework/provider/log"
	"github.com/ddh-open/gin/framework/provider/orm"
	"github.com/ddh-open/gin/framework/provider/redis"
	"github.com/ddh-open/gin/framework/provider/trace"
)

// @title Swagger Devops API
// @version 0.0.1
// @description This is Devops API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name bearer
// @BasePath /
func main() {
	// 初始化服务容器
	container := framework.NewNiceContainer()
	// 绑定App服务提供者
	container.Bind(&app.NiceAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.NiceEnvProvider{})
	container.Bind(&config.NiceConfigProvider{})
	container.Bind(&id.NiceIDProvider{})
	container.Bind(&trace.NiceTraceProvider{})
	container.Bind(&log.NiceLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.ProviderRedis{})
	container.Bind(&cache.NiceCacheProvider{})
	// 业务服务
	container.Bind(&grpc.ProviderGrpc{})
	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := app2.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.NiceKernelProvider{HttpEngine: engine})
	}
	boot.InitService(container)
}
