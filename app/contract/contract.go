package contract

import (
	"google.golang.org/grpc"
)

const KeyGrpc = "devops:grpc"

type ServiceGrpc interface {
	GetGrpc(configPath string, opt ...interface{}) (*grpc.ClientConn, error)
}
