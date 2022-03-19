package path

import (
	"context"
	"encoding/json"
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/sys/model/path"
	"github.com/ddh-open/gin/framework"
	contract2 "github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/resources/proto/userGrpc"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strings"
)

type Service struct {
	repository *base.Repository
}

func NewService(c framework.Container) *Service {
	db, err := c.MustMake(contract2.ORMKey).(contract2.ORMService).GetDB()
	logger := c.MustMake(contract2.LogKey).(contract2.Log)
	if err != nil {
		logger.Error("service 获取db出错： err", zap.Error(err))
	}
	return &Service{base.NewRepository(db, &path.DevopsSysApi{})}
}

func (s *Service) GetRepository() *base.Repository {
	return s.repository
}

func (s *Service) SetRepository(model interface{}) *base.Repository {
	return s.repository.SetRepository(model)
}

func (s *Service) GetApiById(id string, grpcService contract.ServiceGrpc) ([]map[string]interface{}, error) {
	conn, err := grpcService.GetGrpc("grpc.user")
	var result []map[string]interface{}
	if err != nil {
		return result, err
	}
	defer conn.Close()
	client := userGrpc.NewServiceApiClient(conn)
	resp, err := client.ApiList(context.Background(), &userGrpc.ListRequest{
		Filter: []string{"id = ?", id},
	})
	if err != nil {
		return result, err
	}
	if resp.GetResult().GetCode() == 200 {
		err = json.Unmarshal(resp.GetList(), &result)
	} else {
		err = errors.Wrap(err, resp.GetResult().GetMsg())
	}
	return result, err
}

func (s *Service) GetApiList(request base.PageRequest, grpcService contract.ServiceGrpc, param ...interface{}) (base.PageResult, error) {
	conn, err := grpcService.GetGrpc("grpc.user")
	var result base.PageResult
	var list []map[string]interface{}
	if err != nil {
		return result, err
	}
	defer conn.Close()
	client := userGrpc.NewServiceApiClient(conn)
	resp, err := client.ApiList(context.Background(), &userGrpc.ListRequest{
		Filter:   request.Filter,
		Page:     request.Page,
		PageSize: request.PageSize,
	})
	if err != nil {
		return result, err
	}
	if resp.GetResult().GetCode() == 200 {
		err = json.Unmarshal(resp.GetList(), &list)
	} else {
		err = errors.Wrap(err, resp.GetResult().GetMsg())
	}
	result.List = list
	result.PageSize = resp.GetPageSize()
	result.Page = resp.GetPageSize()
	result.Total = resp.GetCounts()
	return result, err
}

func (s *Service) AddApi(mapData map[string]interface{}, grpcService contract.ServiceGrpc, param ...interface{}) error {
	conn, err := grpcService.GetGrpc("grpc.user")
	if err != nil {
		return err
	}
	defer conn.Close()
	data, err := json.Marshal(&mapData)
	if err != nil {
		return err
	}
	client := userGrpc.NewServiceApiClient(conn)
	resp, err := client.ApiAdd(context.Background(), &userGrpc.BytesRequest{
		Data: data,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != 200 {
		err = errors.Wrap(err, resp.GetMsg())
	}
	return err
}

func (s *Service) ModifyApi(mapData map[string]interface{}, grpcService contract.ServiceGrpc, param ...interface{}) error {
	conn, err := grpcService.GetGrpc("grpc.user")
	if err != nil {
		return err
	}
	defer conn.Close()
	data, err := json.Marshal(&mapData)
	if err != nil {
		return err
	}
	client := userGrpc.NewServiceApiClient(conn)
	resp, err := client.ApiModify(context.Background(), &userGrpc.BytesRequest{
		Data: data,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != 200 {
		err = errors.Wrap(err, resp.GetMsg())
	}
	return err
}

func (s *Service) DeleteApi(ids string, grpcService contract.ServiceGrpc, param ...interface{}) error {
	var idsInt []int64
	if strings.Contains(ids, ",") {
		for _, s2 := range strings.Split(ids, ",") {
			idsInt = append(idsInt, cast.ToInt64(s2))
		}
	} else {
		idsInt = append(idsInt, cast.ToInt64(ids))
	}
	conn, err := grpcService.GetGrpc("grpc.user")
	if err != nil {
		return err
	}
	defer conn.Close()
	client := userGrpc.NewServiceApiClient(conn)
	resp, err := client.ApiDelete(context.Background(), &userGrpc.IdsRequest{
		Ids: idsInt,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != 200 {
		err = errors.Wrap(err, resp.GetMsg())
	}
	return err
}
