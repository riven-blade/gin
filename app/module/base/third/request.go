package third

import "github.com/ddh-open/gin/resources/proto/thirdGrpc"

type TencentResourceListRequest struct {
	Type   thirdGrpc.ThirdTencentResourceType `json:"type,omitempty"`
	Ids    []string                           `json:"ids,omitempty"`
	Names  []string                           `json:"names,omitempty"`
	Offset int64                              `json:"offset,omitempty"`
	Limit  int64                              `json:"limit,omitempty"`
}

// TopicListRequest cls
type TopicListRequest struct {
	Offset int64 `json:"offset"` // 分页的偏移量
	Limit  int64 `json:"limit"`  // 分页单页限制数目
}

type TopicCreateRequest struct {
	LogsetId  string `json:"logset_id"`  // 日志集
	TopicName string `json:"topic_name"` // 日志名称
}

type TopicDeleteRequest struct {
	TopicId string `json:"topic_id"` // topicID
}

type DeleteMerchantLog struct {
	MerchantName string `json:"merchant_name"` // 商户名称
	MerchantId   string `json:"merchant_id"`   // 商户id
}

// AddMerchantApmRequest apm
type AddMerchantApmRequest struct {
	MerchantName string   `json:"merchant_name"` // 商户名称
	MerchantId   string   `json:"merchant_id"`   // 商户id
	Namespaces   []string `json:"namespaces"`    // 商户环境
}
