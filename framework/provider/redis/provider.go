package redis

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
)

// ProviderRedis 提供App的具体实现方法
type ProviderRedis struct {
}

// Register 注册方法
func (h *ProviderRedis) Register(container framework.Container) framework.NewInstance {
	return NewNiceRedis
}

// Boot 启动调用
func (h *ProviderRedis) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *ProviderRedis) IsDefer() bool {
	return true
}

// Params 获取初始化参数
func (h *ProviderRedis) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (h *ProviderRedis) Name() string {
	return contract.RedisKey
}
