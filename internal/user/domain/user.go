package domain

// User 用于 service 和 repository 之间的数据交流
type User struct {
	Uid       int64
	UserName  string
	Password  string
	Email     string
	Phone     string
	AvatarURL string
}

// 也可以考虑将业务逻辑封装至 domain 的实体类中
func (u *User) IsValidEmail() bool {
	return true
}

func (u *User) IsValidPhone() bool {
	return true
}
