package pack

import (
	"github.com/schwarzsail/tiktok/internal/user/service/domain"
	"github.com/schwarzsail/tiktok/pkg/kitex_gen/model"
)

// BuildUserInfo 将 domain 中的数据转换成 model 中的数据
func BuildUserInfo(u *domain.User) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		Username: u.UserName,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}
