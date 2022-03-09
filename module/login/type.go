package login

// UserName
// 用户账号登录
type UserName struct {
	// 用户名
	Phone string `json:"phone" binding:"required"`
	// 用户密码
	Password string `json:"password" binding:"required"`
}

// UserCode
// 用户验证码登录
type UserCode struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

// ForgetPassword
// 用户忘记密码
type ForgetPassword struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

// GetPhoneCode
// 获取用户验证码
type GetPhoneCode struct {
	Code string `json:"code" binding:"required"`
}

// User
// 用户登录成功
type User struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Uid   string `json:"id"`
}

// SetPassword
// 用户修改密码操作
type SetPassword struct {
	NewPassword string `json:"new_password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}
