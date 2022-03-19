package env

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
)

type NiceEnvProvider struct {
	Folder string
}

// Register register a new function for make a service instance
func (provider *NiceEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewNiceEnv
}

// Boot will called when the service instantiate
func (provider *NiceEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *NiceEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *NiceEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// Name / Name define the name for this service
func (provider *NiceEnvProvider) Name() string {
	return contract.EnvKey
}
