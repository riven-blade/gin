// Package gin Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MDT style
// license that can be found in the LDCENSE file.
package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// DResponse DResponse代表返回方法
type DResponse interface {
	// DJson Json输出
	DJson(obj interface{}) DResponse

	// DJsonp Jsonp输出
	DJsonp(obj interface{}) DResponse

	// DXml xml输出
	DXml(obj interface{}) DResponse

	// DHtml html输出
	DHtml(template string, obj interface{}) DResponse

	// DText string
	DText(format string, values ...interface{}) DResponse

	// DRedirect 重定向
	DRedirect(path string) DResponse

	// DSetHeader header
	DSetHeader(key string, val string) DResponse

	// DSetCookie Cookie
	DSetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) DResponse

	// DSetStatus 设置状态码
	DSetStatus(code int) DResponse

	// DSetOkStatus 设置200状态
	DSetOkStatus() DResponse
}

// DJsonp Jsonp输出
func (c *Context) DJsonp(obj interface{}) DResponse {
	// 获取请求参数callback
	callbackFunc := c.Query("callback")
	c.DSetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := c.Writer.Write([]byte(callback))
	if err != nil {
		return c
	}
	// 输出左括号
	_, err = c.Writer.Write([]byte("("))
	if err != nil {
		return c
	}
	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return c
	}
	_, err = c.Writer.Write(ret)
	if err != nil {
		return c
	}
	// 输出右括号
	_, err = c.Writer.Write([]byte(")"))
	if err != nil {
		return c
	}
	return c
}

// DXml xml输出
func (c *Context) DXml(obj interface{}) DResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return c.DSetStatus(http.StatusInternalServerError)
	}
	c.DSetHeader("Content-Type", "application/html")
	c.Writer.Write(byt)
	return c
}

// DHtml html输出
func (c *Context) DHtml(file string, obj interface{}) DResponse {
	// 读取模版文件，创建template实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return c
	}
	// 执行Execute方法将obj和模版进行结合
	if err := t.Execute(c.Writer, obj); err != nil {
		return c
	}

	c.DSetHeader("Content-Type", "application/html")
	return c
}

// DText string
func (c *Context) DText(format string, values ...interface{}) DResponse {
	out := fmt.Sprintf(format, values...)
	c.DSetHeader("Content-Type", "application/text")
	c.Writer.Write([]byte(out))
	return c
}

// DRedirect 重定向
func (c *Context) DRedirect(path string) DResponse {
	http.Redirect(c.Writer, c.Request, path, http.StatusMovedPermanently)
	return c
}

// DSetHeader header
func (c *Context) DSetHeader(key string, val string) DResponse {
	c.Writer.Header().Add(key, val)
	return c
}

// DSetCookie Cookie
func (c *Context) DSetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) DResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return c
}

// DSetStatus 设置状态码
func (c *Context) DSetStatus(code int) DResponse {
	c.Writer.WriteHeader(code)
	return c
}

// DSetOkStatus 设置200状态
func (c *Context) DSetOkStatus() DResponse {
	c.Writer.WriteHeader(http.StatusOK)
	return c
}

func (c *Context) DJson(obj interface{}) DResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return c.DSetStatus(http.StatusInternalServerError)
	}
	c.DSetHeader("Content-Type", "application/json")
	c.Writer.Write(byt)
	return c
}
