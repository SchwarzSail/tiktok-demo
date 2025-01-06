package db

import (
	"context"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/model"
	"gorm.io/gorm"
)

// User 数据库所操作的表
type User struct {
	gorm.Model
}

type DBAdapter struct {
	gorm *gorm.DB
}

func NewDBAdapter() *DBAdapter {
	return &DBAdapter{}
}

func (d *DBAdapter) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return nil, nil
}

func (d *DBAdapter) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}
