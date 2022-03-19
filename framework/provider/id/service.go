package id

import (
	"github.com/rs/xid"
)

type NiceIDService struct {
}

func NewNiceIDService(params ...interface{}) (interface{}, error) {
	return &NiceIDService{}, nil
}

func (s *NiceIDService) NewID() string {
	return xid.New().String()
}
