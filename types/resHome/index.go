package resHome

import (
	"modTest/module"
	"modTest/module/user"
)

type MessageUpdate struct {
	Message user.Message `json:"message"`
	Res     string       `json:"res"`
}

type KeysRes struct {
	module.Resp
	Body struct {
		Private string `json:"private"` // 私有密钥
		Public  string `json:"public"`  // 共有密钥
	} `json:"body"`
}

type MessageRes struct {
	module.Resp
	Body user.Message
}

type MessageUpdateRes struct {
	module.Resp
	Body user.Message `json:"body"`
}
