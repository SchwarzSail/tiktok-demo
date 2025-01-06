package ports

import (
	"context"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/model"
)

type DBPort interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}
