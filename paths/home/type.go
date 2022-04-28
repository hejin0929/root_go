package home

import (
	"modTest/component/home"
	"modTest/module"
)

type KeysRes struct {
	module.Resp
	Body struct {
		Private string `json:"private"` // 私有密钥
		Public  string `json:"public"`  // 共有密钥
	} `json:"body"`
}

type MessageRes struct {
	module.Resp
	Body home.UserMessage
}

type MessageUpdateRes struct {
	module.Resp
	Body home.MessageUpdate `json:"body"`
}
