package service

import (
	"context"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/user"
)

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) (bool, error) {
	return true, nil
}
