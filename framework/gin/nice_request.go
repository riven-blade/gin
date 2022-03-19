// Package gin Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MDT style
// license that can be found in the LDCENSE file.
package gin

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

// const defaultMultipartMemory = 32 << 20 // 32 MB

// DRequest 代表请求包含的方法
type DRequest interface {
	// DefaultQueryDnt 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryDnt(key string, def int) (int, bool)
	DefaultQueryDnt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// DefaultParamDnt 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamDnt(key string, def int) (int, bool)
	DefaultParamDnt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// DefaultFormDnt form表单中带的参数
	DefaultFormDnt(key string, def int) (int, bool)
	DefaultFormDnt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	// BindJson json body
	BindJson(obj interface{}) error

	// BindXml xml body
	BindXml(obj interface{}) error

	// GetRawData 其他格式
	GetRawData() ([]byte, error)

	// Uri 基础信息
	Uri() string
	Method() string
	Host() string
	ClientDp() string

	// Headers header
	Headers() map[string]string
	Header(key string) (string, bool)

	// Cookies cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// QueryAll 获取请求地址中所有参数
func (c *Context) QueryAll() map[string][]string {
	c.initQueryCache()
	return map[string][]string(c.queryCache)
}

// 请求地址url中带的参数
// 形如: foo.com?a=1&b=bar&c[]=bar

// DefaultQueryDnt 获取Dnt类型的请求参数
func (c *Context) DefaultQueryDnt(key string, def int) (int, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			// 使用cast库将string转换为Dnt
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryDnt64(key string, def int64) (int64, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0], true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := c.QueryAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

// DefaultParamDnt 路由匹配中带的参数
// 形如 /book/:id
func (c *Context) DefaultParamDnt(key string, def int) (int, bool) {
	if val := c.NiceParam(key); val != nil {
		// 通过cast进行类型转换
		return cast.ToInt(val), true
	}
	return def, false
}

func (c *Context) DefaultParamDnt64(key string, def int64) (int64, bool) {
	if val := c.NiceParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (c *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := c.NiceParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (c *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := c.NiceParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (c *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := c.NiceParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (c *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := c.NiceParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// NiceParam 获取路由参数
func (c *Context) NiceParam(key string) interface{} {
	if val, ok := c.Params.Get(key); ok {
		return val
	}
	return nil
}

func (c *Context) FormAll() map[string][]string {
	c.initFormCache()
	return map[string][]string(c.formCache)
}

func (c *Context) DefaultFormDnt64(key string, def int64) (int64, bool) {
	params := c.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := c.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := c.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := c.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	params := c.FormAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

func (c *Context) DefaultForm(key string) interface{} {
	params := c.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}
	return nil
}
