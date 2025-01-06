package ports

import (
	"context"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/user"
)

type ServicePort interface {
	Register(ctx context.Context, req *user.RegisterRequest) (bool, error)
	Login(ctx context.Context, req *user.LoginRequest) (bool, error)
}
