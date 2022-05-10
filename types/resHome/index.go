package resHome

import (
	"modTest/module"
<<<<<<< HEAD:types/resHome/index.go
	"modTest/module/user"
=======
	user2 "modTest/types/user"
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051:paths/home/type.go
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
<<<<<<< HEAD:types/resHome/index.go
	Body user.Message
=======
	Body user2.Message
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051:paths/home/type.go
}

type MessageUpdateRes struct {
	module.Resp
<<<<<<< HEAD:types/resHome/index.go
	Body user.Message `json:"body"`
=======
	Body user2.MessageUpdate `json:"body"`
>>>>>>> 7c5cfbafc3b65ab7e1446ab7bf96cee5a60d0051:paths/home/type.go
}
