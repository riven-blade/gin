package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/ddh-open/gin/framework"
	contract2 "github.com/ddh-open/gin/framework/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

type Service struct {
	container framework.Container
}

func NewService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &Service{container: container}, nil
}

func (s *Service) GetGrpc(configPath string, opt ...interface{}) (*grpc.ClientConn, error) {
	cert, err := tls.LoadX509KeyPair("./resources/tls/client.pem", "./resources/tls/client.key")
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./resources/tls/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	config := s.container.MustMake(contract2.ConfigKey).(contract2.Config)
	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   config.GetString(configPath + ".serviceName"),
		RootCAs:      certPool,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range opt {
		if token, ok := v.(credentials.PerRPCCredentials); ok {
			return grpc.Dial(config.GetString(configPath+".port"), grpc.WithTransportCredentials(cred), grpc.WithPerRPCCredentials(token))
		}
	}
	return grpc.Dial(config.GetString(configPath+".port"), grpc.WithTransportCredentials(cred))
}
