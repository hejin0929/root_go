package getCode

import "modTest/module"

// GetPhoneCode
// 获取用户验证码
type GetPhoneCode struct {
	Code string `json:"code" binding:"required"`
}

type ResCode struct {
	module.Resp
	Body GetPhoneCode
}
