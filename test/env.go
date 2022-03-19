package test

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/provider/app"
	"github.com/ddh-open/gin/framework/provider/config"
	"github.com/ddh-open/gin/framework/provider/env"
)

const (
	BasePath = "/Users/freemud/Desktop/devops-grpc/"
)

func InitBaseContainer() *framework.NiceContainer {
	// 初始化服务容器
	container := framework.NewNiceContainer()
	// 绑定App服务提供者
	container.Bind(&app.NiceAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.NiceEnvProvider{})
	container.Bind(&config.NiceConfigProvider{})
	return container
}
