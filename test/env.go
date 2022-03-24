package test

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/provider/app"
	"os"
	"strings"
)

var BasePath string

func init() {
	baseFolder, err := os.Getwd()
	if err == nil {
		if strings.Contains(baseFolder, "/framework/") {
			BasePath = strings.Split(baseFolder, "framework")[0]
		} else if strings.Contains(baseFolder, "/app/") {
			BasePath = strings.Split(baseFolder, "app")[0]
		}
	}
}

func InitBaseContainer() *framework.NiceContainer {
	// 初始化服务容器
	container := framework.NewNiceContainer()
	// 绑定App服务提供者
	container.Bind(&app.NiceAppProvider{BaseFolder: BasePath})
	return container
}
