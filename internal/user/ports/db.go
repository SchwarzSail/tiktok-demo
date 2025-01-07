package ports

import (
	"context"
	"github.com/schwarzsail/tiktok/internal/user/service/domain"
)

// DBPort 定义数据库接口，暴露给 core（service）使用
// 最简单的 CURD
type DBPort interface {
	GetUserByName(ctx context.Context, name string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}
