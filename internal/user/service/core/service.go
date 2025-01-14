package service

import "github.com/schwarzsail/tiktok/internal/user/ports"

// CoreService 核心业务，实现 ports 定义的接口
type CoreService struct {
	db ports.DBPort
	// rpc client
}

// NewCoreService 创建核心业务
func NewCoreService(db ports.DBPort) *CoreService {
	return &CoreService{
		db: db,
	}
}

// 在这里定义其他的逻辑业务方法
func (s *CoreService) isEmailValid(email string) bool {
	return true
}

func (s *CoreService) isPhoneValid(phone string) bool {
	return true
}

func (s *CoreService) encryptPassword(password string) string {
	return password
}
