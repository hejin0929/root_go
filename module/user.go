package module

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

// ActiveUserLogin
// 判断用户登陆在线的接口
type ActiveUserLogin struct {
	Model
	Phone string `json:"phone"`
	Login bool   `json:"login"`
	Token string `json:"token"`
}
