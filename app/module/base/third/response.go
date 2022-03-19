package third

import "github.com/ddh-open/gin/app/module/third/model/cls"

type TopicCreateResponse struct {
	TopicId   string `json:"topic_id"` // topicID
	RequestId string `json:"request_id"`
}

type TopicDeleteResponse struct {
	RequestId string `json:"request_id"`
}

type TopicModifyRequest struct {
	TopicId   string `json:"topic_id"`   // topicID
	TopicName string `json:"topic_name"` // 日志名称
	Period    int64  `json:"period"`     // 生命周期
}

type TopicModifyResponse struct {
	RequestId string `json:"request_id"`
}

type AddMerchantClsLogTopicRequest struct {
	MerchantName string   `json:"merchant_name"` // 商户名称
	MerchantId   string   `json:"merchant_id"`   // 商户id
	Namespaces   []string `json:"namespaces"`    // 商户环境
}

type AddMerchantClsLogTopicResponse struct {
	MerchantName string       `json:"merchant_name"` // 商户名称
	MerchantId   string       `json:"merchant_id"`   // 商户id
	Topics       []*cls.Topic `json:"topics"`        // 日志主题列表
}
