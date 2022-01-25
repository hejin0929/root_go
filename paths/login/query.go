package loginPaths

// UserName
// 用户账号登录
type UserName struct {
	// 用户名
	Phone string `json:"phone" binding:"required"`
	// 用户密码
	Password string `json:"password" binding:"required"`
}
