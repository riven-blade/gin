package trace

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
)

type NiceTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *NiceTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewNiceTraceService
}

// Boot will called when the service instantiate
func (provider *NiceTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *NiceTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *NiceTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

// Name / Name define the name for this service
func (provider *NiceTraceProvider) Name() string {
	return contract.TraceKey
}
