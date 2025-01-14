package rpc

import (
	"context"

	"github.com/schwarzsail/tiktok/internal/user/domain"
	"github.com/schwarzsail/tiktok/internal/user/pack"
	"github.com/schwarzsail/tiktok/internal/user/ports"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/user"
)

// UserHandler 实现 idl 定义的 rpc 接口
type UserHandler struct {
	corePort ports.CorePort // 核心业务接口
}

func NewUserHandler(corePort ports.CorePort) *UserHandler {
	return &UserHandler{
		corePort: corePort,
	}
}

func (s *UserHandler) Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	// 判断参数有效性

	// 调用核心业务接口
	u, err := s.corePort.RegisterUser(ctx, &domain.User{
		UserName: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		// 错误处理
	}
	// 调用 pack，整理数据并返回
	r.User = pack.BuildUserInfo(u)
	return
}

func (s *UserHandler) Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	return nil, nil
}
