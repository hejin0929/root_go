package home

import (
	"modTest/module"
	user2 "modTest/types/user"
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
	Body user2.Message
}

type MessageUpdateRes struct {
	module.Resp
	Body user2.MessageUpdate `json:"body"`
}
