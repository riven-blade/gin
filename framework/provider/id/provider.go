package id

import (
	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
)

type NiceIDProvider struct {
}

// Register register a new function for make a service instance
func (provider *NiceIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewNiceIDService
}

// Boot will called when the service instantiate
func (provider *NiceIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *NiceIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *NiceIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

// Name / Name define the name for this service
func (provider *NiceIDProvider) Name() string {
	return contract.IDKey
}
