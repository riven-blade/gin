package cls

type Topic struct {
	LogsetId  string `json:"logset_id"`  // 日志集
	TopicId   string `json:"topic_id"`   // 日志主题id
	TopicName string `json:"topic_name"` // 日志主题名称
}
