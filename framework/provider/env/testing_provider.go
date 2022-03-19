package env

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
)

type NiceTestingEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *NiceTestingEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewNiceTestingEnv
}

// Boot will called when the service instantiate
func (provider *NiceTestingEnvProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *NiceTestingEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *NiceTestingEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

// Name / Name define the name for this service
func (provider *NiceTestingEnvProvider) Name() string {
	return contract.EnvKey
}
