package user

import (
	"context"
	"encoding/json"
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/app/module/base"
	"github.com/ddh-open/gin/app/module/base/sys"
	"github.com/ddh-open/gin/resources/proto/userGrpc"
	"github.com/pkg/errors"
)

func (s *Service) GetRolesByUserId(id string, domain string, grpcService contract.ServiceGrpc, param ...interface{}) ([]map[string]interface{}, error) {
	conn, err := grpcService.GetGrpc("grpc.user")
	var result []map[string]interface{}
	if err != nil {
		return result, err
	}
	defer conn.Close()
	client := userGrpc.NewServiceCabinClient(conn)
	resp, err := client.GetCabinById(context.Background(), &userGrpc.WithSliderParamRequest{
		PType:      "g",
		FieldIndex: 0,
		FieldValue: []string{id, "", domain},
	})
	if err != nil {
		return result, err
	}
	if resp.GetResult().GetCode() == 200 {
		err = json.Unmarshal(resp.GetData(), &result)
	} else {
		err = errors.Wrap(err, resp.GetResult().GetMsg())
	}
	return result, err
}

func (s *Service) RelativeRolesToUser(request sys.RelativeUserRequest, grpcService contract.ServiceGrpc, param ...interface{}) error {
	conn, err := grpcService.GetGrpc("grpc.user")
	if err != nil {
		return err
	}
	defer conn.Close()
	requestData := make([]base.CabinInReceive, 0)
	for _, v := range request.TargetId {
		requestData = append(requestData, base.CabinInReceive{
			PType:    "g",
			Source:   request.UUID,
			Resource: v,
			Domain:   request.Domain,
		})
	}
	requestJson, err := json.Marshal(&requestData)
	if err != nil {
		return err
	}
	client := userGrpc.NewServiceCabinClient(conn)
	resp, err := client.CabinRuleAdd(context.Background(), &userGrpc.BytesRequest{
		Data: requestJson,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != 200 {
		err = errors.Wrap(err, resp.GetMsg())
	}
	return err
}

func (s *Service) DeleteRelativeRolesToUser(request sys.RelativeUserRequest, grpcService contract.ServiceGrpc, param ...interface{}) error {
	conn, err := grpcService.GetGrpc("grpc.user")
	if err != nil {
		return err
	}
	defer conn.Close()
	requestData := make([]base.CabinInReceive, 0)
	for _, v := range request.TargetId {
		requestData = append(requestData, base.CabinInReceive{
			PType:    "g",
			Source:   request.UUID,
			Resource: v,
			Domain:   request.Domain,
		})
	}
	requestJson, err := json.Marshal(&requestData)
	if err != nil {
		return err
	}
	client := userGrpc.NewServiceCabinClient(conn)
	resp, err := client.CabinRuleDelete(context.Background(), &userGrpc.BytesRequest{
		Data: requestJson,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != 200 {
		err = errors.Wrap(err, resp.GetMsg())
	}
	return err
}
