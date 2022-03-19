package user

import (
	"context"
	"encoding/json"
	"github.com/ddh-open/gin/app/contract"
	"github.com/ddh-open/gin/resources/proto/userGrpc"
	"github.com/pkg/errors"
)

func (s *Service) GetMenusByUserId(uuid string, domain string, grpcService contract.ServiceGrpc, param ...interface{}) ([]map[string]interface{}, error) {
	conn, err := grpcService.GetGrpc("grpc.user")
	var result []map[string]interface{}
	if err != nil {
		return result, err
	}
	defer conn.Close()
	client := userGrpc.NewServiceCabinClient(conn)
	resp, err := client.GetCabinById(context.Background(), &userGrpc.WithSliderParamRequest{
		PType:      "g3",
		FieldIndex: 0,
		FieldValue: []string{uuid, "", domain},
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
