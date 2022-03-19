package config

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
	"path/filepath"
)

type NiceConfigProvider struct{}

// Register register a new function for make a service instance
func (provider *NiceConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewNiceConfig
}

// Boot will called when the service instantiate
func (provider *NiceConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *NiceConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *NiceConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

// Name / Name define the name for this service
func (provider *NiceConfigProvider) Name() string {
	return contract.ConfigKey
}
