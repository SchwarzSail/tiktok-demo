package ports

import (
	"context"
	"github.com/schwarzsail/tiktok/internal/user/domain"
)

// CorePort 定义核心业务接口，暴露给 rpc（handler 层）使用
type CorePort interface {
	RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error)
	Login(ctx context.Context, user *domain.User) (*domain.User, error)
}
