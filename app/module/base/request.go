package base

type PageRequest struct {
	Page     int64    `json:"page"`
	PageSize int64    `json:"pageSize"`
	Filter   []string `json:"filter"`
}

// CabinInReceive Cabin structure for input parameters
type CabinInReceive struct {
	PType    string `json:"pType"`    //
	Source   string `json:"source"`   //
	Resource string `json:"resource"` // 路径
	Domain   string `json:"domain"`   //
	Method   string `json:"method"`   //
}
