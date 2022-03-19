package sys

import "github.com/ddh-open/gin/resources/proto/userGrpc"

type LoginRequest struct {
	Username string
	Password string
	Type     userGrpc.UserLoginType
}

type RelativeUserRequest struct {
	UUID     string   `json:"uuid"`
	TargetId []string `json:"targetId"`
	Domain   string   `json:"domain"`
}
