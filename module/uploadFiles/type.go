package uploadFiles

import "modTest/module"

type FilesRes struct {
	module.Resp
	Body string `json:"body"`
}
