package rpc

import (
	"context"
	"github.com/schwarzsail/tiktok/internal/user/ports"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/user"
)

// UserServerImpl 实现了 user rpc 接口
type UserServerImpl struct {
	srv ports.ServicePort
}

func NewUserServerImpl(srv ports.ServicePort) *UserServerImpl {
	return &UserServerImpl{srv: srv}
}

func (s *UserServerImpl) Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	return &user.RegisterResponse{}, nil
}

func (*UserServerImpl) Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	return &user.LoginResponse{}, nil
}
