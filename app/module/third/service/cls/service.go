package cls

import (
	"context"
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/third"
	"github.com/ddh-open/gin/app/module/third/model/cls"
	"github.com/ddh-open/gin/framework"
	contract2 "github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/resources/proto/thirdGrpc"
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
		logger.Error("service 获取db出错： err", zap.Error(err))
	}
	return &Service{base.NewRepository(db, nil)}
}

func (s *Service) GetRepository() *base.Repository {
	return s.repository
}

func (s *Service) SetRepository(model interface{}) *base.Repository {
	return s.repository.SetRepository(model)
}

func (s *Service) DeleteMerchantLog(request third.DeleteMerchantLog, grpcService contract.ServiceGrpc, param ...interface{}) (result base.Response, err error) {
	// 1. 创建日志集，得到日志集id
	conn, err := grpcService.GetGrpc("grpc.third")
	defer conn.Close()
	client := thirdGrpc.NewClsServiceClient(conn)
	if err != nil {
		return result, err
	}

	resp, err := client.DeleteMerchantLog(context.Background(), &thirdGrpc.DeleteMerchantLogRequest{
		MerchantName: request.MerchantName,
		MerchantId:   request.MerchantId,
	})
	if err != nil {
		return result, err
	}
	if resp.GetResult().GetCode() != 200 {
		err = errors.Wrap(err, resp.GetResult().GetMsg())
	}

	result.Msg = resp.GetResult().Msg

	return result, err

}

func (s *Service) AddMerchantClsLogTopic(request third.AddMerchantClsLogTopicRequest, grpcService contract.ServiceGrpc, param ...interface{}) (result base.Response, err error) {
	// 1. 创建日志集，得到日志集id
	conn, err := grpcService.GetGrpc("grpc.third")
	defer conn.Close()
	client := thirdGrpc.NewClsServiceClient(conn)
	if err != nil {
		return result, err
	}
	resp, err := client.CreateLogset(context.Background(), &thirdGrpc.LogsetCreateRequest{
		LogsetName: request.MerchantName, // LogsetName == MerchantName
	})
	if err != nil {
		return result, err
	}
	if resp.GetResult().GetCode() != 200 {
		err = errors.Wrap(err, resp.GetResult().GetMsg())
	}

	logsetId := resp.LogsetId // 获取商户的日志id

	// 2. 创建每个名称空间的日志主题，得到日志主题信息
	res, err := client.BatchCreateClsTopic(context.Background(), &thirdGrpc.BatchCreateClsTopicRequest{
		LogsetId:     logsetId,
		TopicName:    request.Namespaces,
		MerchantId:   request.MerchantId,
		MerchantName: request.MerchantName,
	})
	if err != nil {
		return result, err
	}
	if resp.GetResult().GetCode() != 200 {
		err = errors.Wrap(err, resp.GetResult().GetMsg())
	}

	var topics []*cls.Topic
	for _, v := range res.GetTopics() {
		topics = append(topics, &cls.Topic{
			LogsetId:  v.GetLogsetId(),
			TopicId:   v.GetTopicId(),
			TopicName: v.GetTopicName(),
		})
	}
	result.Msg = resp.Result.GetMsg()
	result.Data = third.AddMerchantClsLogTopicResponse{
		MerchantName: request.MerchantName,
		MerchantId:   request.MerchantId,
		Topics:       topics,
	}

	return result, nil
}
