package service

import (
	"context"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/user"
)

func (s *UserService) Login(ctx context.Context, req *user.LoginRequest) (bool, error) {
	return true, nil
}
