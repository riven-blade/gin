// Package gin Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
)

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}

// context 实现container的几个封装

// Make 实现make的封装
func (c *Context) Make(key string) (interface{}, error) {
	return c.container.Make(key)
}

// MustMake 实现mustMake的封装
func (c *Context) MustMake(key string) interface{} {
	return c.container.MustMake(key)
}

// MakeNew 实现makeNew的封装
func (c *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return c.container.MakeNew(key, params)
}
