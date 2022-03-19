package grpc

import (
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/framework"
)

// ProviderGrpc 服务管理
type ProviderGrpc struct {
	c framework.Container
}

func (pg *ProviderGrpc) Name() string {
	return contract.KeyGrpc
}

func (pg *ProviderGrpc) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (pg *ProviderGrpc) IsDefer() bool {
	return false
}

func (pg *ProviderGrpc) Params(c framework.Container) []interface{} {
	return []interface{}{pg.c}
}

func (pg *ProviderGrpc) Boot(c framework.Container) error {
	pg.c = c
	return nil
}
