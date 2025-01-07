package db

import (
	"context"
	"github.com/schwarzsail/tiktok/internal/user/service/domain"
	"gorm.io/gorm"
)

// 数据库实体
type User struct {
	gorm.Model
}

type DBAdapter struct {
	db *gorm.DB
}

// 不太规范，这样写只是为了不报错
func InitDBConnection() *gorm.DB {
	return nil
}

func NewDBAdapter() *DBAdapter {
	return &DBAdapter{
		db: InitDBConnection(),
	}
}

func (d *DBAdapter) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	return nil, nil
}

func (d *DBAdapter) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}
