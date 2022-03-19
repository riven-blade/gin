package gin

import (
	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/resources/proto/userGrpc"
)

var userKey = "claims@user"

func (c *Context) GetUserKey() string {
	return userKey
}

// MustMakeApp 从容器中获取App服务
func (c *Context) MustMakeApp() contract.App {
	return c.MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Context) MustMakeKernel() contract.Kernel {
	return c.MustMake(contract.KernelKey).(contract.Kernel)
}

// MustMakeConfig 从容器中获取配置服务
func (c *Context) MustMakeConfig() contract.Config {
	return c.MustMake(contract.ConfigKey).(contract.Config)
}

// MustMakeLog 从容器中获取日志服务
func (c *Context) MustMakeLog() contract.Log {
	return c.MustMake(contract.LogKey).(contract.Log)
}

// MustGetUser 获取user
func (c *Context) MustGetUser() *userGrpc.BaseUserInfo {
	return c.MustGet(userKey).(*userGrpc.BaseUserInfo)
}

// GetUser 获取user
func (c *Context) GetUser() (*userGrpc.BaseUserInfo, bool) {
	if user, ok := c.Get(userKey); ok {
		return user.(*userGrpc.BaseUserInfo), ok
	}
	return nil, false
}
