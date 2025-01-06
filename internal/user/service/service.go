package service

import "github.com/schwarzsail/tiktok/internal/user/ports"

type UserService struct {
	db ports.DBPort
}

// NewUserService creates a new UserService, injecting the given DBPort.
func NewUserService(db ports.DBPort) *UserService {
	return &UserService{db: db}
}
