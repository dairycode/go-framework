package service

import (
	"context"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New)

type Service struct{}

func New() (s *Service, cf func(), err error) {
	s = &Service{}
	cf = func() {
		// todo
	}
	return
}

func (s *Service) Ping(ctx context.Context) error {
	// todo 检查 dao 连通性
	return nil
}
