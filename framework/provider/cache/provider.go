package cache

import (
	"strings"

	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/provider/cache/services"
)

// NiceCacheProvider 服务提供者
type NiceCacheProvider struct {
	framework.ServiceProvider

	Driver string // Driver
}

// Register 注册一个服务实例
func (l *NiceCacheProvider) Register(c framework.Container) framework.NewInstance {
	if l.Driver == "" {
		tcs, err := c.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用console
			return services.NewMemoryCache
		}

		cs := tcs.(contract.Config)
		l.Driver = strings.ToLower(cs.GetString("cache.driver"))
	}

	// 根据driver的配置项确定
	switch l.Driver {
	case "redis":
		return services.NewRedisCache
	case "memory":
		return services.NewMemoryCache
	default:
		return services.NewMemoryCache
	}
}

// Boot 启动的时候注入
func (l *NiceCacheProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer 是否延迟加载
func (l *NiceCacheProvider) IsDefer() bool {
	return true
}

// Params 定义要传递给实例化方法的参数
func (l *NiceCacheProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

// Name 定义对应的服务字符串凭证
func (l *NiceCacheProvider) Name() string {
	return contract.CacheKey
}
