package service

import (
	"context"
	"fmt"
	"github.com/schwarzsail/tiktok/internal/user/domain"
)

func (s *CoreService) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// 判断用户是否存在
	u, err := s.db.GetUserByName(ctx, user.UserName)
	if err != nil {
		// 错误上抛，这边应该是数据库的问题
		return nil, fmt.Errorf("get user by name failed: %w", err)
	}
	if u != nil {
		// 用户已存在
		return nil, fmt.Errorf("user already exists")
	}
	// 判断邮箱和手机号是否有效（被注册过）
	if !s.isPhoneValid(user.Phone) || !s.isEmailValid(user.Email) {
		return nil, fmt.Errorf("phone or email is invalid")
	}

	// 密码加密
	user.Password = s.encryptPassword(user.Password)
	// 存入数据库
	u, err = s.db.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create user failed: %w", err)
	}
	// 返回用户信息
	return u, nil
}
