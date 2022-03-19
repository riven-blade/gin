package kernel

import (
	"github.com/ddh-open/gin/framework/gin"
	"net/http"
)

// NiceKernelService 引擎服务
type NiceKernelService struct {
	engine *gin.Engine
}

// NewNiceKernelService 初始化web引擎服务实例
func NewNiceKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &NiceKernelService{engine: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *NiceKernelService) HttpEngine() http.Handler {
	return s.engine
}
