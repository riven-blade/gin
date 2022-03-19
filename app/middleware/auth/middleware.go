package auth

import (
	"context"
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/framework/gin"
	"github.com/ddh-open/gin/resources/proto/userGrpc"
)

var whiteList = append(make([]string, 0), "/user/login", "/user/register")

// MiddlewareAuth 登录中间件
func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("bearer")
		authPass := false
		for i := range whiteList {
			if c.Request.RequestURI == whiteList[i] {
				authPass = true
			}
		}
		// debug 模式开启， 跳过权限验证
		if c.MustMakeConfig().GetBool("app.debug") {
			authPass = true
		}
		// 不在白名单， 通过grpc鉴权
		if !authPass {
			grpcToken := base.Token{Value: token}
			devops := c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc)
			conn, err := devops.GetGrpc("grpc.user", &grpcToken)
			if err != nil {
				base.FailWithDetailed(gin.H{"reload": true}, "用户api鉴权连接出错:"+err.Error(), c)
				c.Abort()
				return
			}
			defer conn.Close()
			client := userGrpc.NewServiceAuthClient(conn)
			resp, err := client.AuthApi(context.Background(), &userGrpc.AuthApiRequest{
				Path: c.Request.RequestURI,
			})
			if err != nil {
				base.FailWithDetailed(gin.H{"reload": true}, "用户api鉴权调用出错:"+err.Error(), c)
				c.Abort()
				return
			}
			// 代表响应成功
			if resp.GetResult().GetCode() != 200 {
				base.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问:"+resp.GetResult().GetMsg(), c)
				c.Abort()
				return
			}
			c.Set(c.GetUserKey(), resp.GetUser())
		}
		c.Next()
	}
}
