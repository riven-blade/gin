package user

import (
	"context"
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/sys"
	"github.com/ddh-open/gin/app/module/sys/model/user"
	"github.com/ddh-open/gin/framework"
	contract2 "github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/gin"
	"github.com/ddh-open/gin/resources/proto/userGrpc"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service struct {
	repository *base.Repository
}

func NewService(c framework.Container) *Service {
	db, err := c.MustMake(contract2.ORMKey).(contract2.ORMService).GetDB()
	logger := c.MustMake(contract2.LogKey).(contract2.Log)
	if err != nil {
		logger.Error("Service 获取db出错： err", zap.Error(err))
	}
	return &Service{base.NewRepository(db, &user.SysUser{})}
}

func (s *Service) GetRepository() *base.Repository {
	return s.repository
}

func (s *Service) GetUsers() {

}

func (s *Service) Login(request sys.LoginRequest, grpc contract.ServiceGrpc) (interface{}, error) {
	result := make(map[string]string, 1)
	conn, err := grpc.GetGrpc("grpc.user")
	if err != nil {
		err = errors.Wrap(err, "初始化grpc连接出错")
		return result, err
	}
	defer conn.Close()
	client := userGrpc.NewUserServiceClient(conn)
	resp, err := client.Login(context.Background(), &userGrpc.WithPasswordRequest{
		Username: request.Username,
		Password: request.Password,
		Type:     request.Type,
	})
	if err != nil {
		err = errors.Wrap(err, "grpc 登录接口出错")
		return result, err
	}
	// 代表响应成功
	if resp.GetResult().Code != 200 {
		err = errors.Wrap(errors.New("grpc code -1"), resp.GetResult().GetMsg())
	}
	result["token"] = resp.GetToken()
	return result, err
}

// GetUserInfo 获取用户详细信息
func (s *Service) GetUserInfo(c *gin.Context) (interface{}, error) {
	devops := c.MustMake(contract.KeyGrpc).(contract.ServiceGrpc)
	conn, err := devops.GetGrpc("grpc.user")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := userGrpc.NewUserServiceClient(conn)
	client.GetUserInfo(context.Background(), &userGrpc.GetUserInfoRequest{
		UserId: 1,
	})
	return nil, nil
}
